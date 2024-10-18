package scheduler

import (
	"context"
	"edgecom.ai/timeseries/internal/services"
	"edgecom.ai/timeseries/pkg/models"
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
				log.Println("Scheduler Started...")
				s.runDataFetching()
			}
		}
	}()
}

func (s *scheduler) runDataFetching() {
	start, err := s.timeSeriesService.GetLatestDataPointTimestamp(context.Background())
	if err != nil {
		log.Printf("Error fetching latest data point timestamp: %v", err)
		return
	}

	end := time.Now()
	dataCh := make(chan models.TimeSeriesResult)
	defer close(dataCh)

	go func() {
		for data := range dataCh {
			err := s.scraper.StoreData(data)
			if err != nil {
				log.Printf("Error storing data: %v", err)
			}
		}
	}()

	err = s.scraper.FetchData(context.Background(), time.Unix(start, 0), end, dataCh)
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		return
	}
	log.Println("Data scheduler finished successfully.")
}
