package bootstrap

import (
	"context"
	"edgecom.ai/timeseries/internal/services"
	"edgecom.ai/timeseries/pkg/models"
	"log"
	"sync"
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

	var wg sync.WaitGroup
	stop := time.Now()
	start := stop.AddDate(-2, 0, 0)
	chunkSize := 24 * time.Hour

	ch := make(chan []models.TimeSeriesData)
	defer close(ch)

	go func() {
		for data := range ch {
			b.service.StoreData(data)
		}
	}()

	for i := start; i.Before(stop); i = i.Add(chunkSize) {
		wg.Add(1)
		startChunk := i
		endChunk := i.Add(chunkSize)

		if endChunk.After(stop) {
			endChunk = stop
		}
		go func(start, end time.Time) {
			defer wg.Done()
			if err := b.service.FetchData(context.Background(), start, end, ch); err != nil {
				log.Printf("Error fetching data: %v", err)
			}
		}(startChunk, endChunk)
	}

	wg.Wait()
	log.Println("Bootstrap completed successfully.")
}
