package app

import (
	"edgecom.ai/timeseries/internal/repository/influxdb"
	"edgecom.ai/timeseries/internal/services"
)

func (a *application) InitService(r influxdb.Repository) services.TimeSeriesService {
	return services.NewTimeSeriesService(r)
}
