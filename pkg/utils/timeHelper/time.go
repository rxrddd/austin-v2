package timeHelper

import "time"

const (
	DateYMD           = "20060102"
	DateDefaultLayout = "2006-01-02 15:04:05"
)

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

// FormatTimeYMDHIS 获取指定时间戳的年月日时分秒
func FormatTimeYMDHIS(timestamp time.Time) string {
	return timestamp.Format("2006-01-02 15:04:05")
}

func FormatTimeInt64YMDHIS(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

// GetDisTodayEnd 获取现在到今天结束的秒数
func GetDisTodayEnd() int64 {
	todayEnd, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	todayEndUnix := todayEnd.AddDate(0, 0, 1).Unix()
	period := todayEndUnix - time.Now().Unix()
	return period
}
