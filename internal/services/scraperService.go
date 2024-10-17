package services

import (
	"context"
	"edgecom.ai/timeseries/internal/repository/influxdb"
	"edgecom.ai/timeseries/pkg/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type ScraperService interface {
	FetchData(start, end time.Time, dataCh chan []models.TimeSeriesData) error
	StoreData(data []models.TimeSeriesData)
}

type scraperService struct {
	client     *http.Client
	repository influxdb.Repository
	baseUrl    string
}

type ResponseData struct {
	Result []models.TimeSeriesData `json:"result"`
}

func NewScraperService(r influxdb.Repository, baseUrl string) ScraperService {
	return &scraperService{&http.Client{}, r, baseUrl}
}

func (s *scraperService) FetchData(start, end time.Time, dataCh chan []models.TimeSeriesData) error {
	timeFormat := "2006-01-02T15:04:05"
	u := fmt.Sprintf("%s?start=%s&end=%s", s.baseUrl, start.Format(timeFormat), end.Format(timeFormat))

	log.Printf("Fetching time series data from API ... [%s - %s]", start.Format(timeFormat), end.Format(timeFormat))

	response, err := s.client.Get(u)
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

func (s *scraperService) StoreData(data []models.TimeSeriesData) {
	if err := s.repository.WriteData(context.Background(), data); err != nil {
		log.Printf("Error writing data: %v", err)
	}
}
