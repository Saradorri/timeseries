package services

import (
	"context"
	"edgecom.ai/timeseries/internal/repository"
	"edgecom.ai/timeseries/pkg/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type ScraperService interface {
	FetchData(ctx context.Context, start, end time.Time, dataCh chan models.TimeSeriesResult) error
	StoreData(ctx context.Context, data models.TimeSeriesResult) error
}

type scraperService struct {
	client     *http.Client
	repository repository.Repository
	baseUrl    string
}

type ResponseData struct {
	Result models.TimeSeriesResult `json:"result"`
}

func NewScraperService(r repository.Repository, baseUrl string) ScraperService {
	return &scraperService{&http.Client{}, r, baseUrl}
}

func (s *scraperService) FetchData(ctx context.Context, start, end time.Time, dataCh chan models.TimeSeriesResult) error {
	timeFormat := "2006-01-02T15:04:05"
	u := fmt.Sprintf("%s?start=%s&end=%s", s.baseUrl, start.Format(timeFormat), end.Format(timeFormat))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	response, err := s.client.Do(req)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return fmt.Errorf("fetch data timeout: %w", err)
		}
		return fmt.Errorf("failed to fetch data: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("failed to close response body")
		}
	}(response.Body)

	var apiResponse ResponseData
	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	log.Printf("Fetched time series data from API ... [%s - %s]", start.Format(timeFormat), end.Format(timeFormat))

	select {
	case dataCh <- apiResponse.Result:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("context canceled while sending data: %w", ctx.Err())
	}
}
func (s *scraperService) StoreData(ctx context.Context, data models.TimeSeriesResult) error {
	if err := s.repository.WriteData(ctx, data); err != nil {
		return fmt.Errorf("failed to store data: %w", err)
	}
	return nil
}
