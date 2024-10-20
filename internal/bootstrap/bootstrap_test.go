package bootstrap_test

import (
	"context"
	mocks "edgecom.ai/timeseries/mocks/services"
	"errors"
	"fmt"
	"testing"
	"time"

	"edgecom.ai/timeseries/internal/bootstrap"
	"edgecom.ai/timeseries/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInitializeHistoricalData_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockScraperService(ctrl)
	b := bootstrap.NewBootstrap("https://example.com", mockService)

	mockService.EXPECT().FetchData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, start, end time.Time, ch chan models.TimeSeriesResult) error {
			ch <- models.TimeSeriesResult{} // Simulate sending data
			return nil
		}).AnyTimes()

	mockService.EXPECT().StoreData(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := b.InitializeHistoricalData(ctx)
	assert.NoError(t, err)
}

func TestInitializeHistoricalData_FetchDataError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockScraperService(ctrl)
	b := bootstrap.NewBootstrap("https://example.com", mockService)

	mockService.EXPECT().FetchData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(fmt.Errorf("fetch data error")).AnyTimes()

	mockService.EXPECT().StoreData(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := b.InitializeHistoricalData(ctx)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error while initializing historical data")
}

func TestInitializeHistoricalData_StoreDataError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockScraperService(ctrl)

	mockService.EXPECT().FetchData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, start, end time.Time, ch chan models.TimeSeriesResult) error {
			ch <- models.TimeSeriesResult{}
			return nil
		}).AnyTimes()

	mockService.EXPECT().StoreData(gomock.Any(), gomock.Any()).Return(errors.New("store error")).Times(1)

	b := bootstrap.NewBootstrap("https://example.com", mockService)

	err := b.InitializeHistoricalData(context.Background())
	assert.Error(t, err)
}

func TestInitializeHistoricalData_ContextDone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockScraperService(ctrl)

	mockService.EXPECT().FetchData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	b := bootstrap.NewBootstrap("https://example.com", mockService)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := b.InitializeHistoricalData(ctx)
	assert.Error(t, err)
}
