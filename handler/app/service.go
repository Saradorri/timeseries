package app

import (
	"edgecom.ai/timeseries/internal/repository/prometheus"
	"edgecom.ai/timeseries/internal/services"
)

func (a *application) InitService(r prometheus.Repository) services.TimeSeriesService {
	return services.NewTimeSeriesService(r)
}
