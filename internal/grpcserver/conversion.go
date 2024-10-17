package grpcserver

import (
	tpb "edgecom.ai/timeseries/internal/proto/pb"
	"edgecom.ai/timeseries/pkg/models"
)

func toProtoResponse(results []models.TimeSeriesData, agr, win string) *tpb.QueryResponse {
	response := &tpb.QueryResponse{
		Meta: &tpb.QueryMetadata{
			Aggregation: agr,
			Window:      win,
			Status:      tpb.QueryStatus_ERROR,
		},
	}

	for _, result := range results {
		tsData := &tpb.TimeSeriesData{
			Time:  result.Time,
			Value: result.Value,
		}
		response.Data = append(response.Data, tsData)
	}

	if len(results) > 0 {
		response.Meta.Status = tpb.QueryStatus_SUCCESS
		response.Meta.Message = "Query executed successfully"
	} else {
		response.Meta.Status = tpb.QueryStatus_ERROR
		response.Meta.Message = "No data found for the given parameters"
	}

	return response
}
