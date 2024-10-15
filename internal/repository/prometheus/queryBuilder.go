package prometheus

import (
	"fmt"
	"strings"
)

var supportedAggregations = map[string]string{
	"MIN": "min_over_time",
	"MAX": "max_over_time",
	"AVG": "avg_over_time",
	"SUM": "sum_over_time",
}

func buildPromQL(metric, aggregation, window string) (string, error) {
	agr := strings.ToUpper(aggregation)

	aggFunc, ok := supportedAggregations[agr]
	if !ok {
		return "", fmt.Errorf("unsupported aggregation: %v", agr)
	}

	promQL := fmt.Sprintf("%s(%s[%s])", aggFunc, metric, window)
	return promQL, nil
}
