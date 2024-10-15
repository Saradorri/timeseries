package prometheus

import (
	"context"
	"edgecom.ai/timeseries/pkg/models"
	"edgecom.ai/timeseries/utils"
	"fmt"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"log"
	"time"
)

type Repository interface {
	Query(ctx context.Context, query models.TimeSeriesQuery) ([]models.TimeSeriesData, error)
}

type repository struct {
	client v1.API
}

func NewRepository(client v1.API) Repository {
	return &repository{client: client}
}

func (r *repository) Query(ctx context.Context, query models.TimeSeriesQuery) ([]models.TimeSeriesData, error) {
	promQL, err := buildPromQL("custom_metric", query.Aggregation, query.Window)
	if err != nil {
		return nil, fmt.Errorf("error building PromQL: %v", err)
	}

	step, err := utils.ParseWindow(query.Window)
	if err != nil {
		log.Fatalf("Error parsing window: %v", err)
	}

	result, warnings, err := r.client.QueryRange(ctx, promQL, v1.Range{
		Start: time.Unix(query.Start, 0),
		End:   time.Unix(query.End, 0),
		Step:  step,
	})
	if err != nil {
		return nil, err
	}

	if len(warnings) > 0 {
		log.Printf("Warnings during query: %v\n", warnings)
	}

	data := convertResultToTimeSeries(result)
	return data, nil
}
