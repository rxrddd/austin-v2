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

// 检查日期格式
func CheckDateFormat(date string) bool {
	format := "2006-01-02"
	dateA, err := time.Parse(format, date)
	if err != nil {
		return false
	}
	dateB := dateA.Format("2006-01-02")
	if dateB != date {
		return false
	}
	return true
}

// 检查时间格式
func CheckDateTimeFormat(dateTime string) bool {
	format := "2006-01-02 15:04:05"
	dateA, err := time.Parse(format, dateTime)
	if err != nil {
		return false
	}
	dateB := dateA.Format("2006-01-02 15:04:05")
	if dateB != dateTime {
		return false
	}
	return true
}
