package goutil

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// FormatTime 格式化时间格式，生成一个描述性的时间
func FormatTime(t time.Time) string {
	now := time.Now()
	duration := now.Sub(t)

	var timeDesc string
	if duration.Seconds() < 60 {
		timeDesc = "刚刚"
	} else if duration.Minutes() < 60 {
		t, _ := ToInt(duration.Minutes())
		timeDesc = fmt.Sprintf("%d 分钟前", t)
	} else if duration.Hours() < 24 {
		t, _ := ToInt(duration.Hours())
		timeDesc = fmt.Sprintf("%d 小时前", t)
	} else {
		timeDesc = now.Format("2006-01-02 15:04")
	}

	return timeDesc
}

// FormatTimeUnix 格式化时间格式
func FormatTimeUnix(t int64) string {
	return FormatTime(time.Unix(t, 0))
}

// MonthCount 计算年月的天数
// @year
// @month
func MonthCount(year int, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30
		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}

	return days
}

// JudgeSmallThenToday 判断是否为小于今天的时间
// @time 要判断的时间
func JudgeSmallThenToday(t int64) bool {
	if GetTimeZero(time.Now()) > t {
		return true
	}
	return false
}

// GetTimeZero 获取指定日期的0点的时间
func GetTimeZero(t time.Time) int64 {
	year, month, day := t.Date()
	yearStr := strconv.Itoa(year)
	dayStr := strconv.Itoa(day)
	nowt, _ := time.Parse("2006-January-2", yearStr+"-"+month.String()+"-"+dayStr)
	return nowt.Unix()
}

// GetTimeByDuration 根据间隔获取指定时间
func GetTimeByDuration(days float64) time.Time {
	hours := int(math.Ceil(days * 24))
	durationStr := fmt.Sprintf("%vh", hours)
	d, _ := time.ParseDuration(durationStr)
	return time.Now().Add(d)
}
