package go_time_utils

import (
    "time"
)
// DateTimeUtils 获取周几，几月 数字和英文
type  DateTimeUtils struct{
    Time time.Time
}
var TimeLayout = "2006-01-02 15:04:05"
var TimeLayoutDate = "2006-01-02"
var TimeLayoutZh = "2006年01月02日 15:04:05"
var TimeLayoutZhDate = "2006年01月02日"
var TimeLayoutZhDateNo = "2006年1月2日"
// FirstDayWeekDay 一年第一天是周几
func (dt *DateTimeUtils ) FirstDayWeekDay() (weekday int) {
    yday := dt.Time.YearDay()
    // 新年第一天
    yFirstDayTime := dt.Time.AddDate(0, 0, -yday + 1)
    // 元旦是周几
    weekday = int(yFirstDayTime.Weekday())
    if weekday == 0 {
        weekday = 7
    }
    return
}
// Weekth 获取时间是第几周
func (dt *DateTimeUtils ) Weekth() (weekth int) {
    yday := dt.Time.YearDay()
    // 新年第一天
    yFirstDayTime := dt.Time.AddDate(0, 0, -yday + 1)
    // 元旦是周几
    firstDayWeekDay := int(yFirstDayTime.Weekday())
    println("yday:", yday, "firstDayWeekDay", firstDayWeekDay)
    // 第一周有几天
    firstWeekDays := 1
    if firstDayWeekDay != 0 {
        firstWeekDays = 7 - firstDayWeekDay + 1
    }
    weekth = 1
    if yday > firstWeekDays {
        weekth = (yday - firstWeekDays)/7 + 1
    }
    return
}
func (dt *DateTimeUtils) Ymd() (int,int,int) {

    y, m1, d := dt.Time.Date()
    m := int(m1)
    return y, m, d
}
func (dt *DateTimeUtils) Month() (m int, s string) {
    month := dt.Time.Month()
    return int(month), month.String()
}
// Dayth 获取时间是一年第几天
func (dt *DateTimeUtils ) Dayth() int {

    return dt.Time.YearDay()
}
// Weekday 获取指定日期是周几
func (dt *DateTimeUtils)Weekday() (weekday int) {
    wday := dt.Time.Weekday()
    weekday = int(wday)
    if weekday == 0 {
        weekday = 7
    }
    return
}
// PassedWeekdayCount 计算已过去多少个周几
func (dt *DateTimeUtils)PassedWeekdayCount(weekday int) (counter int) {
    dtWeekday := dt.Weekday()
    weekth := dt.Weekth()
    counter = weekth
    if dtWeekday < weekday {
        counter = weekth - 1
    }

    // 下面要累计或排除新年第一周的数据
    firstDayWeekDay := dt.FirstDayWeekDay()
    if firstDayWeekDay > weekday {
        counter = counter - 1
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
// StrToTime "2021-06-29 12:00:01"
func StrToTime(dateStr string) time.Time {
    //ts := time.Now().Format(timeLayout)
    tt, _ := time.Parse(TimeLayout, dateStr)
    return tt
}
// StrToTimeDate "2021-06-29"
func StrToTimeDate(dateStr string) time.Time {
    tt, _ := time.Parse(TimeLayoutDate, dateStr)
    return tt
}

// StrToTimeZh "2021年06月29日 12:00:01"
func StrToTimeZh(dateStr string) time.Time {
    tt, _ := time.Parse(TimeLayoutZh, dateStr)
    return tt
}

func StrToTimeZhDate(dateStr string) time.Time {
    tt, _ := time.Parse(TimeLayoutZhDate, dateStr)
    return tt
}
func StrToTimeZhDateNo(dateStr string) time.Time {
    tt, _ := time.Parse(TimeLayoutZhDateNo, dateStr)
    return tt
}