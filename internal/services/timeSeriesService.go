package services

import (
	"context"
	"edgecom.ai/timeseries/internal/repository"
	"edgecom.ai/timeseries/pkg/models"
)

type TimeSeriesService interface {
	GetByQuery(ctx context.Context, q models.TimeSeriesQuery) (models.TimeSeriesResult, error)
	GetLatestDataPointTimestamp(ctx context.Context) (int64, error)
}

type timeSeriesService struct {
	repository repository.Repository
}

func NewTimeSeriesService(r repository.Repository) TimeSeriesService {
	return &timeSeriesService{
		repository: r,
	}
}

func (ts *timeSeriesService) GetByQuery(ctx context.Context, q models.TimeSeriesQuery) (models.TimeSeriesResult, error) {
	return ts.repository.QueryData(ctx, q)
}

func (ts *timeSeriesService) GetLatestDataPointTimestamp(ctx context.Context) (int64, error) {
	return ts.repository.GetLatestDataPointTimestamp(ctx)
}
