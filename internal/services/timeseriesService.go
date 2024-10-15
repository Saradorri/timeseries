package services

import (
	"edgecom.ai/timeseries/internal/metrics"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type TimeSeriesService interface {
	FetchData(endpoint string)
}

type timeSeriesService struct {
	client *http.Client
}

type UnixTimeStamp int64

type TimeSeriesData struct {
	Timestamp int64   `json:"time"`
	Value     float64 `json:"value"`
}

type responseData struct {
	Result []TimeSeriesData `json:"result"`
}

func NewTimeSeriesService() TimeSeriesService {
	return &timeSeriesService{
		&http.Client{},
	}
}

func (s *timeSeriesService) FetchData(endpoint string) {
	log.Println("Fetching time series data from API...")

	response, err := s.client.Get(endpoint)
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error reading data: %v", err)
		}
	}(response.Body)

	var apiResponse responseData

	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		log.Printf("Error decoding API response: %v", err)
		return
	}

	for _, entry := range apiResponse.Result {
		metrics.SetValue(strconv.FormatInt(entry.Timestamp, 10), entry.Value)
	}
}
