package gouse

import (
	"strconv"
	"time"
)

func Second() int {
	return time.Now().Second()
}

func Minute() int {
	return time.Now().Minute()
}

func Hour() int {
	return time.Now().Hour()
}

func Day() int {
	return time.Now().Day()
}

func Month() int {
	return int(time.Now().Month())
}

func Year() int {
	return time.Now().Year()
}

func Weekday() int {
	return int(time.Now().Weekday())
}

func Unix() int64 {
	return time.Now().Unix()
}

func UnixMilli() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func UnixMicro() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

func UnixNano() int64 {
	return time.Now().UnixNano()
}

func UnixMilliToTime(milli int64) time.Time {
	return time.Unix(0, milli*int64(time.Millisecond))
}

func UnixMicroToTime(micro int64) time.Time {
	return time.Unix(0, micro*int64(time.Microsecond))
}

func UnixNanoToTime(nano int64) time.Time {
	return time.Unix(0, nano)
}

func ToSecond(second int) time.Duration {
	return time.Duration(second) * time.Second
}

func ToMinute(minute int) time.Duration {
	return time.Duration(minute) * time.Minute
}

func ToHour(hour int) time.Duration {
	return time.Duration(hour) * time.Hour
}

func SleepSecond(second int) {
	time.Sleep(ToSecond(second))
}

func SleepMinute(minute int) {
	time.Sleep(ToMinute(minute))
}

func SleepHour(hour int) {
	time.Sleep(ToHour(hour))
}

/* Clock format */

func formatTime(t time.Time) string {
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()

	hourStr := strconv.Itoa(hour)
	minuteStr := strconv.Itoa(minute)
	secondStr := strconv.Itoa(second)

	if hour < 10 {
		hourStr = Sprintf("0%s", hourStr)
	}
	if minute < 10 {
		minuteStr = Sprintf("0%s", minuteStr)
	}
	if second < 10 {
		secondStr = Sprintf("0%s", secondStr)
	}

	return Sprintf("%s:%s:%s", hourStr, minuteStr, secondStr)
}

func TerminalClock() {
	msgTime := make(chan time.Time)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			msgTime <- time.Now()
		}
	}()

	for t := range msgTime {
		Cls()
		Println(formatTime(t))
	}
}

func DiffTime(t1, t2 time.Time) time.Duration {
	return t2.Sub(t1)
}

func DiffTimeNow(startTime time.Time) time.Duration {
	elapsedTime := float64(time.Since(startTime).Seconds() * 1000)
	return time.Duration(elapsedTime)
}
