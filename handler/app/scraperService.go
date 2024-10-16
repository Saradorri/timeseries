package app

import "edgecom.ai/timeseries/internal/services"

func (a *application) InitScraper() services.TimeSeriesScraperService {
	return services.NewTimeSeriesScraperService()
}
