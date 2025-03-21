package gouse

import "time"

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

func Milli() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func Micro() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

func Nano() int64 {
	return time.Now().UnixNano()
}

func Milli2Time(milli int64) time.Time {
	return time.Unix(0, milli*int64(time.Millisecond))
}

func Micro2Time(micro int64) time.Time {
	return time.Unix(0, micro*int64(time.Microsecond))
}

func Nano2Time(nano int64) time.Time {
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
		Println(Time2Str(t))
	}
}

func DiffTime(t1, t2 time.Time) time.Duration {
	return t2.Sub(t1)
}

func DiffTimeNow(startTime time.Time) time.Duration {
	elapsedTime := float64(time.Since(startTime).Seconds() * 1000)
	return time.Duration(elapsedTime)
}
