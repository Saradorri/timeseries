package app

import (
	"edgecom.ai/timeseries/internal/repository"
	"edgecom.ai/timeseries/internal/services"
)

func (a *application) InitService(r repository.Repository) services.TimeSeriesService {
	return services.NewTimeSeriesService(r)
}
