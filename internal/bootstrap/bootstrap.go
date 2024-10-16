package bootstrap

import (
	"edgecom.ai/timeseries/internal/services"
	"log"
	"time"
)

type Bootstrap interface {
	InitializeHistoricalData()
}

type bootstrap struct {
	url     string
	service services.ScraperService
}

func NewBootstrap(url string, service services.ScraperService) Bootstrap {
	return &bootstrap{
		url:     url,
		service: service,
	}
}

func (b *bootstrap) InitializeHistoricalData() {
	log.Println("Bootstrapping historical data...")

	now := time.Now()
	twoYearsAgo := now.AddDate(-2, 0, 0)

	err := b.service.FetchData(twoYearsAgo, now)
	if err != nil {
		return
	}
}
