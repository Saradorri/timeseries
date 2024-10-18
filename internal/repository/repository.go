package repository

import (
	"context"
	"edgecom.ai/timeseries/pkg/models"
)

type Repository interface {
	WriteData(ctx context.Context, data models.TimeSeriesResult) error
	QueryData(ctx context.Context, query models.TimeSeriesQuery) (models.TimeSeriesResult, error)
	GetLatestDataPointTimestamp(ctx context.Context) (int64, error)
}
