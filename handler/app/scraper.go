package app

import (
	"edgecom.ai/timeseries/internal/repository/influxdb"
	"edgecom.ai/timeseries/internal/services"
)

func (a *application) InitScraper(r influxdb.Repository) services.ScraperService {
	return services.NewScraperService(r)
}
