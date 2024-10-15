package prometheus

import (
	"edgecom.ai/timeseries/pkg/models"
	"github.com/prometheus/common/model"
)

func convertResultToTimeSeries(value model.Value) []models.TimeSeriesData {
	var result []models.TimeSeriesData

	switch v := value.(type) {
	case model.Matrix:
		for _, stream := range v {
			for _, point := range stream.Values {
				dataPoint := models.TimeSeriesData{
					Timestamp: point.Timestamp.Time().Unix(),
					Value:     float64(point.Value),
				}
				result = append(result, dataPoint)
			}
		}
	case model.Vector:
		for _, sample := range v {
			dataPoint := models.TimeSeriesData{
				Timestamp: sample.Timestamp.Time().Unix(),
				Value:     float64(sample.Value),
			}
			result = append(result, dataPoint)
		}
	default:
		return nil
	}

	return result
}
