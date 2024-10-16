package services

import (
	"context"
	"edgecom.ai/timeseries/internal/repository/influxdb"
	"edgecom.ai/timeseries/pkg/models"
)

type TimeSeriesService interface {
	GetByQuery(ctx context.Context, q models.TimeSeriesQuery) ([]models.TimeSeriesData, error)
}

type timeSeriesService struct {
	repository influxdb.Repository
}

func NewTimeSeriesService(r influxdb.Repository) TimeSeriesService {
	return &timeSeriesService{
		repository: r,
	}
}

func (ts *timeSeriesService) GetByQuery(ctx context.Context, q models.TimeSeriesQuery) ([]models.TimeSeriesData, error) {
	return ts.repository.QueryData(ctx, q)
}
