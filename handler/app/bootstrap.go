package app

import (
	"edgecom.ai/timeseries/internal/bootstrap"
	"edgecom.ai/timeseries/internal/services"
)

func (a *application) InitBootstrap(service services.TimeSeriesScraperService) bootstrap.Bootstrap {
	return bootstrap.NewBootstrap(a.config.App.ApiUrl, service)
}
