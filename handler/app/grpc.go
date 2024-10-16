package app

import (
	"edgecom.ai/timeseries/internal/grpcserver"
	"edgecom.ai/timeseries/internal/services"
)

func (a *application) InitServer(ts services.TimeSeriesService) grpcserver.GrpcServer {
	return grpcserver.NewServer(a.config.App.GrpcPort, ts)
}
