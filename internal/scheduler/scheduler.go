package scheduler

import (
	"context"
	"edgecom.ai/timeseries/internal/services"
	"log"
	"time"
)

type Scheduler interface {
	StartScheduler()
}

type scheduler struct {
	ticker            *time.Ticker
	scraper           services.ScraperService
	timeSeriesService services.TimeSeriesService
}

func NewScheduler(interval int, ss services.ScraperService, ts services.TimeSeriesService) Scheduler {
	ticker := time.NewTicker(time.Duration(interval) * time.Minute)
	return &scheduler{
		ticker:            ticker,
		scraper:           ss,
		timeSeriesService: ts,
	}
}

func (s *scheduler) StartScheduler() {
	go func() {
		for {
			select {
			case <-s.ticker.C:
				log.Println("Fetching new data...")
				start, err := s.timeSeriesService.GetLatestDataPointTimestamp(context.Background())
				err = s.scraper.FetchData(time.Unix(start, 0), time.Now())
				if err != nil {
					return
				}
			}
		}
	}()
}
