package app

import (
	"context"
	"edgecom.ai/timeseries/internal/bootstrap"
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
			a.InitServices,
		),
		fx.Invoke(func(bootstrap bootstrap.Bootstrap) {
			bootstrap.InitializeHistoricalData()
		}),
	)
	app.Run()
}
