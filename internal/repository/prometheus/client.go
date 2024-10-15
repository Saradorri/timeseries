package prometheus

import (
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"log"
)

type PrometheusClient struct {
	Client v1.API
}

func NewPrometheusClient(prometheusURL string) (*PrometheusClient, error) {
	client, err := api.NewClient(api.Config{
		Address: prometheusURL,
	})
	if err != nil {
		log.Printf("Error creating Prometheus client: %v", err)
		return nil, err
	}

	v1api := v1.NewAPI(client)

	return &PrometheusClient{
		Client: v1api,
	}, nil
}
