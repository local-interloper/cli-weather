package utils

import (
	"time"
)

func IsoToWeekday(isoDate string) string {
	date, err := time.Parse(time.DateOnly, isoDate)
	if err != nil {
		panic(err.Error())
	}

	return date.Weekday().String()
}
