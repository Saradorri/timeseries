package services

import (
	"context"
	"edgecom.ai/timeseries/internal/repository/prometheus"
	"edgecom.ai/timeseries/pkg/models"
)

type TimeSeriesService interface {
	GetByQuery(ctx context.Context, q models.TimeSeriesQuery) ([]models.TimeSeriesData, error)
}

type timeSeriesService struct {
	repository prometheus.Repository
}

func NewTimeSeriesService(r prometheus.Repository) TimeSeriesService {
	return &timeSeriesService{
		repository: r,
	}
}

func (ts *timeSeriesService) GetByQuery(ctx context.Context, q models.TimeSeriesQuery) ([]models.TimeSeriesData, error) {
	return ts.repository.Query(ctx, q)
}
