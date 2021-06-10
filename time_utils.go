package go_time_utils

import (
    "time"
)
// DateTimeUtils 获取周几，几月 数字和英文
type  DateTimeUtils struct{
    time time.Time
}
var TimeLayout = "2006-01-02 15:04:05"
// FirstDayWeekDay 一年第一天是周几
func (dt *DateTimeUtils ) FirstDayWeekDay() (weekday int) {
    yday := dt.time.YearDay()
    // 新年第一天
    yFirstDayTime := dt.time.AddDate(0, 0, -yday + 1)
    // 元旦是周几
    weekday = int(yFirstDayTime.Weekday())
    if weekday == 0 {
        weekday = 7
    }
    return
}
// Weekth 获取时间是第几周
func (dt *DateTimeUtils ) Weekth() (weekth int) {
    yday := dt.time.YearDay()
    // 新年第一天
    yFirstDayTime := dt.time.AddDate(0, 0, -yday + 1)
    // 元旦是周几
    firstDayWeekDay := int(yFirstDayTime.Weekday())
    // 第一周有几天
    firstWeekDays := 1
    if firstDayWeekDay != 0 {
        firstWeekDays = 7 - firstDayWeekDay + 1
    }
    weekth = 1
    if yday > firstWeekDays {
        weekth = (yday - firstWeekDays)/7 + 1 + 1
    }
    return
}
func (dt *DateTimeUtils) Ymd() (int,int,int) {

    y, m1, d := dt.time.Date()
    m := int(m1)
    return y, m, d
}
func (dt *DateTimeUtils) Month() (m int, s string) {
    month := dt.time.Month()
    return int(month), month.String()
}
// Dayth 获取时间是一年第几天
func (dt *DateTimeUtils ) Dayth() int {

    return dt.time.YearDay()
}
// Weekday 获取指定日期是周几
func (dt *DateTimeUtils)Weekday() (weekday int) {
    wday := dt.time.Weekday()
    weekday = int(wday)
    if weekday == 0 {
        weekday = 7
    }
    return
}

// StrTimeDiff 两个时间字符串的时间差
func StrTimeDiff(maxTime string, minTime string) time.Duration {
    max := StrToTime(maxTime)
    min := StrToTime(minTime)
    result := max.Sub(min)
    return result
}
func StrToTime(dateStr string) time.Time {
    //ts := time.Now().Format(timeLayout)
    tt, _ := time.Parse(timeLayout, dateStr)
    return tt
}

