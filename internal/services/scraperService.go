package services

import (
	"edgecom.ai/timeseries/internal/repository/influxdb"
	"edgecom.ai/timeseries/pkg/models"
	"encoding/json"
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
)

type ScraperService interface {
	FetchData(endpoint string) error
}

type scraperService struct {
	client     *http.Client
	repository influxdb.Repository
}

type ResponseData struct {
	Result []models.TimeSeriesData `json:"result"`
}

func NewScraperService(r influxdb.Repository) ScraperService {
	return &scraperService{&http.Client{}, r}
}

func (s *scraperService) FetchData(endpoint string) error {
	log.Println("Fetching time series data from API...")

	response, err := s.client.Get(endpoint)
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error reading data: %v", err)
		}
	}(response.Body)

	var apiResponse ResponseData

	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		log.Printf("Error decoding API response: %v", err)
		return err
	}

	err = s.repository.WriteData(context.Background(), apiResponse.Result)
	if err != nil {
		return err
	}
	return nil
}
