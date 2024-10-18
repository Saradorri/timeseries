package validation

import (
	"edgecom.ai/timeseries/pkg/models"
	"errors"
	"fmt"
	"strings"
)

func ValidateQueryRequest(input models.TimeSeriesQuery) error {
	if input.Start >= input.End {
		return errors.New("start time must be before end time")
	}

	if input.Window == "" {
		return errors.New("window cannot be empty")
	}

	validAggregations := map[string]bool{
		"MIN": true,
		"MAX": true,
		"AVG": true,
		"SUM": true,
	}

	if !validAggregations[strings.ToUpper(input.Aggregation)] {
		return fmt.Errorf("invalid aggregation type: %s", input.Aggregation)
	}

	return nil
}
