package main

import (
	"time"
)

func getCurrentDay24HourTimeStamp() bool {
	isInDelta := false
	t := time.Now()

	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	tm2 := tm1.AddDate(0, 0, 1)
	size := tm2.Unix() - t.Unix()
	if size < 60 {
		isInDelta := true
		return isInDelta
	}

	return isInDelta
}

func GetCurrentDateNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
