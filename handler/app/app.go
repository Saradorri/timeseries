package app

import (
	"context"
	"edgecom.ai/timeseries/utils"
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
}
