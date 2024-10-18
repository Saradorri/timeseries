package services

import (
	"context"
	"edgecom.ai/timeseries/internal/repository"
	"edgecom.ai/timeseries/pkg/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type ScraperService interface {
	FetchData(ctx context.Context, start, end time.Time, dataCh chan models.TimeSeriesResult) error
	StoreData(data models.TimeSeriesResult) error
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

	log.Printf("Fetching time series data from API ... [%s - %s]", start.Format(timeFormat), end.Format(timeFormat))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	response, err := s.client.Do(req)

	if err != nil {
		log.Printf("Error fetching data: %v", err)
		return err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}(response.Body)

	var apiResponse ResponseData
	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		log.Printf("Error decoding API response: %v", err)
		return err
	}

	dataCh <- apiResponse.Result
	return nil
}

func (s *scraperService) StoreData(data models.TimeSeriesResult) error {
	if err := s.repository.WriteData(context.Background(), data); err != nil {
		log.Printf("Error writing data: %v", err)
		return err
	}
	return nil
}
