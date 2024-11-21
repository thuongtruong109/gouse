package gouse

import (
	"fmt"
	"strconv"
	"time"
)

/* Date formats */

func format(input interface{}, format string) string {
	var t time.Time

	switch v := input.(type) {
	case time.Time:
		t = v
	case string:
		var err error
		t, err = time.Parse(time.RFC3339, v)
		if err != nil {
			return ""
		}
	default:
		return ""
	}

	return t.Format(format)
}

func ISODate(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006-01-02T15:04:05.999Z")
	} else {
		return format(date[0], "2006-01-02T15:04:05.999Z")
	}
}

func NormalDate(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "01/02/2006")
	} else {
		return format(date[0].(time.Time), "01/02/2006")
	}
}

func ReverseDate(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006/01/02")
	} else {
		return format(date[0].(time.Time), "2006/01/02")
	}
}

func DashDate(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006-01-02")
	} else {
		return format(date[0].(time.Time), "2006-01-02")
	}
}

func DotDate(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006.01.02")
	} else {
		return format(date[0].(time.Time), "2006.01.02")
	}
}

func UnderlineDate(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006_01_02")
	} else {
		return format(date[0].(time.Time), "2006_01_02")
	}
}

func SpaceDate(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006 01 02")
	} else {
		return format(date[0].(time.Time), "2006 01 02")
	}
}

func MonthDate(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "01/2006")
	} else {
		return format(date[0].(time.Time), "01/2006")
	}
}

func LongDate(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "Monday, January 2, 2006")
	} else {
		return format(date[0].(time.Time), "Monday, January 2, 2006")
	}
}

func UTCDate(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "Jan 2, 2006 at 3:04pm (MST)")
	} else {
		return format(date[0].(time.Time), "Jan 2, 2006 at 3:04pm (MST)")
	}
}

/* Time formats */

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

/* Clock format */

func formatTime(t time.Time) string {
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()

	hourStr := strconv.Itoa(hour)
	minuteStr := strconv.Itoa(minute)
	secondStr := strconv.Itoa(second)

	if hour < 10 {
		hourStr = fmt.Sprintf("0%s", hourStr)
	}
	if minute < 10 {
		minuteStr = fmt.Sprintf("0%s", minuteStr)
	}
	if second < 10 {
		secondStr = fmt.Sprintf("0%s", secondStr)
	}

	return fmt.Sprintf("%s:%s:%s", hourStr, minuteStr, secondStr)
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
		fmt.Println(formatTime(t))
	}
}

/* Utilities */

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
