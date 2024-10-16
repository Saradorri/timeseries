package models

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
