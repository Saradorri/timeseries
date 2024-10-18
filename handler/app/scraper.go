package app

import (
	"edgecom.ai/timeseries/internal/repository"
	"edgecom.ai/timeseries/internal/services"
)

func (a *application) InitScraper(r repository.Repository) services.ScraperService {
	return services.NewScraperService(r, a.config.App.ApiUrl)
}
