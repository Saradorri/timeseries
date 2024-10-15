package app

import "edgecom.ai/timeseries/internal/services"

func (a *application) InitServices() services.TimeSeriesService {
	return services.NewTimeSeriesService()
}
