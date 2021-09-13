package utils

import "time"

func GetCurrentTimePlusSeconds(seconds int) time.Time {
	currentTime := time.Now().UTC()
	return currentTime.Add(time.Second * time.Duration(seconds))
}

func TimeHasPassed(aTime time.Time) bool {
	now := time.Now().UTC()
	return aTime.Before(now)
}
