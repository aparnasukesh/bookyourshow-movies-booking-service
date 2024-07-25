package utils

import (
	"time"
)

func ParseDateString(dateString string) (*time.Time, error) {
	layout := "2006-01-02" // Layout to match the format of the date string

	parsedDate, err := time.Parse(layout, dateString)
	if err != nil {
		return nil, err
	}
	return &parsedDate, nil
}
