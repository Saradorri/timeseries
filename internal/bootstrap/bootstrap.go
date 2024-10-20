package bootstrap

import (
	"context"
	"edgecom.ai/timeseries/internal/services"
	"edgecom.ai/timeseries/pkg/models"
	"fmt"
	"log"
	"sync"
	"time"
)

type Bootstrap interface {
	InitializeHistoricalData(ctx context.Context) error
	Close()
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

func (b *bootstrap) InitializeHistoricalData(ctx context.Context) error {
	log.Println("Bootstrapping historical data...")

	var wg sync.WaitGroup
	errCh := make(chan error, 1)
	stop := time.Now()
	start := stop.AddDate(-2, 0, 0)
	chunkSize := 24 * time.Hour

	ch := make(chan models.TimeSeriesResult)

	// store data to influxDB
	go func() {
		for data := range ch {
			storeCtx, storeCancel := context.WithTimeout(ctx, 60*time.Second)
			defer storeCancel()

			if err := b.service.StoreData(storeCtx, data); err != nil {
				errCh <- err
				return
			}
		}
	}()

	// worker pool and define max for it
	const maxWorkers = 100
	sem := make(chan struct{}, maxWorkers)

	for i := start; i.Before(stop); i = i.Add(chunkSize) {
		wg.Add(1)
		startChunk := i
		endChunk := i.Add(chunkSize)

		if endChunk.After(stop) {
			endChunk = stop
		}

		go func(start, end time.Time) {
			defer wg.Done()

			sem <- struct{}{}        // acquire a worker
			defer func() { <-sem }() // releasee the worker

			fetchCtx, fetchCancel := context.WithTimeout(ctx, 60*time.Second)
			defer fetchCancel()

			if err := b.service.FetchData(fetchCtx, start, end, ch); err != nil {
				errCh <- fmt.Errorf("failed to fetch data: %w", err)
			}
		}(startChunk, endChunk)
	}

	doneCh := make(chan struct{})
	go func() {
		wg.Wait()
		close(ch)
		close(doneCh)
	}()

	select {
	case err := <-errCh:
		return fmt.Errorf("error while initializing historical data: %v", err.Error())
	case <-ctx.Done():
		log.Println("Parent context done!")
		return ctx.Err()
	case <-doneCh:
		log.Println("All historical data fetched and stored successfully.")
		return nil
	}
}

func (b *bootstrap) Close() {
	log.Println("Closing bootstrap data...")
}
