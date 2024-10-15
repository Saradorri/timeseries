package scheduler

import (
	"edgecom.ai/timeseries/internal/services"
	"log"
	"time"
)

type Scheduler interface {
	StartScheduler()
}

type scheduler struct {
	ticker  *time.Ticker
	url     string
	service services.TimeSeriesService
}

func NewScheduler(apiUrl string, interval int, service services.TimeSeriesService) Scheduler {
	ticker := time.NewTicker(time.Duration(interval) * time.Minute)
	return &scheduler{
		ticker:  ticker,
		url:     apiUrl,
		service: service,
	}
}

func (s *scheduler) StartScheduler() {
	go func() {
		for {
			select {
			case <-s.ticker.C:
				log.Println("Fetching new data...")
				s.service.FetchData(s.url)
			}
		}
	}()
}
