package influxdb

import (
	"edgecom.ai/timeseries/pkg/models"
	"fmt"
	"strings"
	"time"
)

var supportedAggregations = map[string]string{
	"MIN": "min",
	"MAX": "max",
	"AVG": "mean",
	"SUM": "sum",
}

func RangeQuery(q models.TimeSeriesQuery, bucket, measurement string) (string, error) {

	agr := strings.ToUpper(q.Aggregation)
	aggFunc, ok := supportedAggregations[agr]

	if !ok {
		return "", fmt.Errorf("unsupported aggregation: %v", agr)
	}

	startTime := time.Unix(q.Start, 0).Format(time.RFC3339)
	endTime := time.Unix(q.End, 0).Format(time.RFC3339)

	query := fmt.Sprintf(`
        from(bucket: "%s")
        |> range(start: %s, stop: %s)
        |> filter(fn: (r) => r._measurement == "%s")
        |> aggregateWindow(every: %s, fn: %s)
        |> yield(name: "result")
    `,
		bucket,
		startTime,
		endTime,
		measurement,
		q.Window,
		aggFunc,
	)

	return query, nil
}

func LatestTimestampQuery(bucket string) string {
	query := fmt.Sprintf(`from(bucket: "%s")
              |> range(start: 0)
              |> last()
	`, bucket)

	return query
}
