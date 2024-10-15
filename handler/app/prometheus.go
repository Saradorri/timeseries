package app

import (
	"edgecom.ai/timeseries/internal/metrics"
	"edgecom.ai/timeseries/internal/repository/prometheus"
	"log"
)

func (a *application) InitPrometheus() prometheus.Repository {
	pc, err := prometheus.NewPrometheusClient(a.config.Database.ScrapeURL)
	if err != nil {
		log.Fatal(err)
	}
	metrics.StartServer(a.config.App.MetricPort)
	return prometheus.NewRepository(pc.Client)
}
