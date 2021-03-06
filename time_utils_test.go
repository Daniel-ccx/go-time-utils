package go_time_utils

import (
    "fmt"
    "runtime"
    "strconv"
    "testing"
)
// 获取正在运行的函数名
func runFuncName()string{
    pc := make([]uintptr,1)
    runtime.Callers(2,pc)
    f := runtime.FuncForPC(pc[0])
    return f.Name()
}
func TestStrToTimeDate(t *testing.T) {
    println(runFuncName())
    dt := StrToTimeDate("2021-06-29")
    fmt.Printf("%v", dt)
}
func TestPassedWeekdayCount(t *testing.T) {
    dt := StrToTime("2021-06-08 10:10:10")
    dtu := &DateTimeUtils{
        Time: dt,
    }
    wday := 2
    passedCount := dtu.PassedWeekdayCount(wday)
    wday4 := 4
    passedCount4 := dtu.PassedWeekdayCount(wday4)
    wday7 := 7
    passedCount7 := dtu.PassedWeekdayCount(wday7)
    println("已过去", passedCount, "个周", wday, passedCount4, passedCount7)
}

func TestFirstDayWeekDay(t *testing.T) {
    println(runFuncName())
    dt := StrToTime("2021-06-09 10:10:10")

    dtu := &DateTimeUtils{
        Time: dt,
    }
    openDays := map[int]int{2:2, 4:4, 7:7}
    println(openDays[dtu.FirstDayWeekDay()], "=tessssss")
    println("2021年第一天是周 ", dtu.FirstDayWeekDay())
}
func TestWeekth(t *testing.T) {
    println(runFuncName())
    dt := StrToTime("2021-01-03 10:10:10")

    _ = &DateTimeUtils{
       dt,
    }
    //println("2021年第 ", dtu.Weekth(), " 周的周 ", dtu.Weekday())
}
func TestWeekday(t *testing.T) {
    println(runFuncName())
    dt := StrToTime("2021-01-03 10:10:10")

    dtu := &DateTimeUtils{
         dt,
    }
    println(dtu.Weekday())
}
func TestStrToTime(t *testing.T) {
    println(runFuncName())
    _ = StrToTime("2020-10-01 10:10:10")
    //println(tt.String())
}
func TestStrTimeDiff(t *testing.T) {
    println(runFuncName())

    _ = StrTimeDiff("2021-10-01 10:10:10", "2021-10-01 00:10:00")

    //println(float64ToStr(sub.Hours()), float64ToStr(sub.Seconds()))
}
func float64ToStr(f float64) string {
    return strconv.FormatFloat(f, 'f', -1, 64)
}