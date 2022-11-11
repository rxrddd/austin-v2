package timeHelper

import "time"

// CurrentTimeYMDHIS 获取年月日时分秒
func CurrentTimeYMDHIS() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// CurrentTimeYMD 获取年月日
func CurrentTimeYMD() string {
	return time.Now().Format("2006-01-02")
}

// CurrentTimeHI 获取时分
func CurrentTimeHI() string {
	return time.Now().Format("15:04")
}

// CurrentTimeH 获取时
func CurrentTimeH() string {
	return time.Now().Format("15")
}

// CurrentTimeHIS 获取时分秒
func CurrentTimeHIS() string {
	return time.Now().Format("15:04:05")
}

// FormatTimeYMGHIS 获取指定时间戳的年月日时分秒
func FormatTimeYMGHIS(timestamp time.Time) string {
	return timestamp.Format("2006-01-02 15:04:05")
}

func FormatTimeInt64YMGHIS(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}
