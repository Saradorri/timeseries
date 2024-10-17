package influxdb

import (
	"context"
	"edgecom.ai/timeseries/pkg/models"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"log"
	"time"
)

type Repository interface {
	WriteData(ctx context.Context, data []models.TimeSeriesData) error
	QueryData(ctx context.Context, query models.TimeSeriesQuery) ([]models.TimeSeriesData, error)
	GetLatestDataPointTimestamp(ctx context.Context) (int64, error)
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
			SetTime(time.Unix(point.Time, 0))

		if err := writeAPI.WritePoint(ctx, p); err != nil {
			return fmt.Errorf("failed to write point to InfluxDB: %w", err)
		}
	}
	log.Println("Finished writing data to InfluxDB")
	return nil
}

func (r *influxDBRepository) QueryData(ctx context.Context, q models.TimeSeriesQuery) ([]models.TimeSeriesData, error) {
	measurement := "time_series_data"
	query, err := RangeQuery(q, r.bucket, measurement)

	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	result, err := r.Client.QueryAPI(r.org).Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	var results []models.TimeSeriesData
	for result.Next() {
		if result.Err() != nil {
			return nil, fmt.Errorf("query error: %v", result.Err())
		}

		timestamp := result.Record().Time().Unix()
		value := result.Record().Value()

		if value == nil {
			log.Printf("Warning: Got nil value for timestamp: %v", timestamp)
			continue
		}

		floatValue, ok := value.(float64)
		if !ok {
			return nil, fmt.Errorf("unexpected value type: %T", value)
		}

		results = append(results, models.TimeSeriesData{
			Time:  timestamp,
			Value: floatValue,
		})
	}
	return results, nil
}

func (r *influxDBRepository) GetLatestDataPointTimestamp(ctx context.Context) (int64, error) {
	query := LatestTimestampQuery(r.bucket)

	result, err := r.Client.QueryAPI(r.org).Query(ctx, query)
	if err != nil || result == nil {
		return 0, err
	}

	var latestTimestamp int64
	for result.Next() {
		record := result.Record()
		latestTimestamp = record.Time().Unix()
	}
	return latestTimestamp, nil
}
