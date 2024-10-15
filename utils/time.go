package utils

import (
	"fmt"
	"strconv"
	"time"
)

func ParseWindow(windowStr string) (time.Duration, error) {
	unit := windowStr[len(windowStr)-1:]
	value := windowStr[:len(windowStr)-1]

	durationValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid duration value: %v", err)
	}

	switch unit {
	case "s":
		return time.Duration(durationValue) * time.Second, nil
	case "m":
		return time.Duration(durationValue) * time.Minute, nil
	case "h":
		return time.Duration(durationValue) * time.Hour, nil
	case "d":
		return time.Duration(durationValue) * 24 * time.Hour, nil
	case "w":
		return time.Duration(durationValue) * 7 * 24 * time.Hour, nil
	default:
		return 0, fmt.Errorf("unsupported time unit: %s", unit)
	}
}
