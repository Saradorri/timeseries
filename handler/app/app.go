package app

import (
	"context"
	"edgecom.ai/timeseries/internal/bootstrap"
	"edgecom.ai/timeseries/internal/grpcserver"
	"edgecom.ai/timeseries/internal/scheduler"
	"edgecom.ai/timeseries/utils"
	"go.uber.org/fx"
	"log"
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
			a.InitPrometheus,
			a.InitBootstrap,
			a.InitScraper,
			a.InitScheduler,
			a.InitServer,
			a.InitService,
		),
		fx.Invoke(func(bootstrap bootstrap.Bootstrap, scheduler scheduler.Scheduler, server grpcserver.GrpcServer) {
			bootstrap.InitializeHistoricalData()
			scheduler.StartScheduler()
			err := server.StartServer()
			if err != nil {
				return
			}
		}),
	)
	app.Run()
}
