package models

import tpb "edgecom.ai/timeseries/internal/proto/pb"

type TimeSeriesQuery struct {
	Start       int64
	End         int64
	Window      string
	Aggregation string
}

type TimeSeriesData struct {
	Time  int64
	Value float64
}

type TimeSeriesResult []TimeSeriesData

func (tr *TimeSeriesResult) ToProtoResponse(agr, win string) *tpb.QueryResponse {
	response := &tpb.QueryResponse{
		Meta: &tpb.QueryMetadata{
			Aggregation: agr,
			Window:      win,
			Status:      tpb.QueryStatus_ERROR,
		},
	}

	for _, result := range *tr {
		tsData := &tpb.TimeSeriesData{
			Time:  result.Time,
			Value: result.Value,
		}
		response.Data = append(response.Data, tsData)
	}

	if len(*tr) > 0 {
		response.Meta.Status = tpb.QueryStatus_SUCCESS
		response.Meta.Message = "Query executed successfully"
	} else {
		response.Meta.Status = tpb.QueryStatus_ERROR
		response.Meta.Message = "No data found for the given parameters"
	}

	return response
}
