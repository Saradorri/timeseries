package bootstrap

import (
	"edgecom.ai/timeseries/internal/services"
	"fmt"
	"log"
	"time"
)

type Bootstrap interface {
	InitializeHistoricalData()
}

type bootstrap struct {
	url     string
	service services.TimeSeriesService
}

func NewBootstrap(url string, service services.TimeSeriesService) Bootstrap {
	return &bootstrap{
		url:     url,
		service: service,
	}
}

func (b *bootstrap) InitializeHistoricalData() {
	log.Println("Bootstrapping historical data...")

	now := time.Now()
	twoYearsAgo := now.AddDate(-2, 0, 0)
	timeFormat := "2006-01-02T15:04:05"

	u := fmt.Sprintf("%s?start=%s&end=%s", b.url, twoYearsAgo.Format(timeFormat), now.Format(timeFormat))

	b.service.FetchData(u)
}
