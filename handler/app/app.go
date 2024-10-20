package app

import (
	"context"
	"edgecom.ai/timeseries/internal/bootstrap"
	"edgecom.ai/timeseries/internal/grpcserver"
	"edgecom.ai/timeseries/internal/scheduler"
	"edgecom.ai/timeseries/utils"
	"fmt"
	"go.uber.org/fx"
	"log"
	"time"
)

type Application interface {
	Setup()
	GetContext() context.Context
}

type application struct {
	ctx    context.Context
	config *utils.ServiceConfig
}

func NewApplication(ctx context.Context) Application {
	return &application{ctx: ctx}
}

func (a *application) GetContext() context.Context {
	return a.ctx
}

func (a *application) Setup() {
	err := a.setupViper()
	if err != nil {
		log.Panic(err.Error())
	}

	app := fx.New(
		fx.Provide(
			a.InitInfluxDB,
			a.InitBootstrap,
			a.InitScraper,
			a.InitScheduler,
			a.InitServer,
			a.InitService,
		),
		fx.Invoke(func(bootstrap bootstrap.Bootstrap, scheduler scheduler.Scheduler, server grpcserver.GrpcServer, lc fx.Lifecycle) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					signal := make(chan error, 1)

					bootstrapCtx, bootstrapCancel := context.WithTimeout(context.Background(), 30*time.Minute) // new context because of long process
					schedulerCtx, schedulerCancel := context.WithCancel(context.Background())                  // without timeout

					go func() {
						log.Println("Starting bootstrap initialization...")
						if err := bootstrap.InitializeHistoricalData(bootstrapCtx); err != nil {
							bootstrapCancel()
							log.Printf("Error initializing bootstrap: %v", err)
							signal <- err
							return
						}
						signal <- nil
					}()

					go func() {
						if err := <-signal; err != nil {
							log.Printf("Bootstrap failed: %v", err.Error())
							schedulerCancel()
							return
						}
						log.Println("Bootstrap completed successfully, starting the scheduler...")
						scheduler.StartScheduler(schedulerCtx)
					}()

					if err := server.StartServer(); err != nil {
						bootstrapCancel()
						schedulerCancel()
						return fmt.Errorf("failed to start gRPC server: %v", err.Error())
					}

					return nil
				},
				OnStop: func(ctx context.Context) error {
					log.Println("Stopping scheduler and closing bootstrap...")
					scheduler.StopScheduler()
					bootstrap.Close()
					return nil
				},
			})
		}),
	)
	app.Run()
}
