package helper

import (
	"fmt"
	"time"
)

func ConvertTimeStringToBeginningOfMounth(date string) (string, error) {
	layout := "01-2006"
	t, err := time.Parse(layout, date)
	if err != nil {
		return "", fmt.Errorf("Error parsing start_date should be format 01-2006")
	}

	return t.Format(time.DateOnly), nil
}

func ConvertTimeStringToTheEndOfMounth(date string) (string, error) {
	layout := "01-2006"
	t, err := time.Parse(layout, date)
	if err != nil {
		return "", fmt.Errorf("Error parsing end_date should be format 01-2006")
	}

	firstDayOfNextMonth := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
	t = firstDayOfNextMonth.Add(-time.Second)

	return t.Format(time.DateOnly), nil
}
