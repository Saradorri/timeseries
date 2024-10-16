package influxdb

import (
	"context"
	"edgecom.ai/timeseries/pkg/models"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

type Repository interface {
	WriteData(ctx context.Context, data []models.TimeSeriesData) error
	QueryData(ctx context.Context, query models.TimeSeriesQuery) ([]models.TimeSeriesData, error)
}

type influxDBRepository struct {
	*InfluxDBClient
	org    string
	bucket string
}

func NewRepository(c *InfluxDBClient, org, bucket string) Repository {
	return &influxDBRepository{c, org, bucket}
}

func (r *influxDBRepository) WriteData(ctx context.Context, data []models.TimeSeriesData) error {
	measurement := "time_series_data"
	writeAPI := r.Client.WriteAPIBlocking(r.org, r.bucket)
	for _, point := range data {
		p := influxdb2.NewPointWithMeasurement(measurement).
			AddTag("source", "api").
			AddField("value", point.Value).
			SetTime(time.Unix(point.Timestamp, 0))

		if err := writeAPI.WritePoint(ctx, p); err != nil {
			return fmt.Errorf("failed to write point to InfluxDB: %w", err)
		}
	}
	return nil
}

func (r *influxDBRepository) QueryData(ctx context.Context, q models.TimeSeriesQuery) ([]models.TimeSeriesData, error) {
	query, err := BuildQuery(q, r.bucket, q.Aggregation)

	result, err := r.Client.QueryAPI(r.org).Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	var results []models.TimeSeriesData
	for result.Next() {
		value := result.Record().Value().(float64)
		timestamp := result.Record().Time().Unix()
		results = append(results, models.TimeSeriesData{
			Timestamp: timestamp,
			Value:     value,
		})
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	return results, nil
}
