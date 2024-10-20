package scheduler

import (
	"context"
	"edgecom.ai/timeseries/internal/services"
	"edgecom.ai/timeseries/pkg/models"
	"fmt"
	"log"
	"sync"
	"time"
)

type Scheduler interface {
	StartScheduler(ctc context.Context)
	StopScheduler()
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

func (s *scheduler) StartScheduler(ctx context.Context) {
	go func() {
		for {
			select {
			case <-s.ticker.C:
				log.Println("Scheduler starting...")
				if err := s.runDataFetching(ctx); err != nil {
					log.Println("Scheduler failed:", err)
					return
				}
			case <-ctx.Done(): // parent context cancellation
				log.Println("Scheduler exit due to parent context cancellation.")
				return
			}
		}
	}()
}

func (s *scheduler) runDataFetching(ctx context.Context) error {
	errCh := make(chan error, 1)
	end := time.Now()
	dataCh := make(chan models.TimeSeriesResult)

	start, err := s.timeSeriesService.GetLatestDataPointTimestamp(ctx)
	if err != nil {
		return fmt.Errorf("failed to get latest data point: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range dataCh {
			select {
			case <-ctx.Done():
				log.Println("StoreData canceled due to context cancellation")
				return
			default:
				if err := s.scraper.StoreData(ctx, data); err != nil {
					errCh <- fmt.Errorf("failed to store data: %v", err)
					return
				}
				log.Println("Scheduler finished.")
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.scraper.FetchData(ctx, time.Unix(start, 0), end, dataCh); err != nil {
			errCh <- fmt.Errorf("failed to fetch data: %v", err)
			return
		}
		close(dataCh)
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	select {
	case err := <-errCh:
		return fmt.Errorf("error in scheduler: %v", err.Error())
	case <-ctx.Done():
		log.Println("fetching data in scheduler stopped due to context cancellation")
		return ctx.Err()
	default:
		return nil
	}
}

func (s *scheduler) StopScheduler() {
	log.Println("Stopping scheduler...")
	s.ticker.Stop()
}
