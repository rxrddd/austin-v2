package timeHelper

import "time"

func GetCurrentYMDHIS() string {
	now := GetCurrentTime()
	return now.Format("2006-01-02 15:04:05")
}

func FormatYMDHIS(timing *time.Time) string {
	return timing.Format("2006-01-02 15:04:05")
}

func GetCurrentTime() *time.Time {
	now := time.Now()
	return &now
}
