package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

func SampleDateTime() {
	println("Second:", gouse.Second())
	println("Minute:", gouse.Minute())
	println("Hour:", gouse.Hour())
	println("Day:", gouse.Day())
	println("Month:", gouse.Month())
	println("Year:", gouse.Year())
	println("Weekday:", gouse.Weekday())
	println("Unix:", gouse.Unix())
	println("UnixNano:", gouse.UnixNano())
	println("UnixMilli:", gouse.UnixMilli())
	println("UnixMicro:", gouse.UnixMicro())
	fmt.Println("UnixMilliToTime:", gouse.UnixMilliToTime(1000000000))
	fmt.Println("UnixMicroToTime:", gouse.UnixMicroToTime(1000000000))
	fmt.Println("UnixNanoToTime:", gouse.UnixNanoToTime(1000000000))
}

func SampleDateISO() {
	println("ISO:", gouse.ISODate())
}

func SampleDateShort() {
	println("ShortNormal:", gouse.NormalDate())
	println("ShortReverse:", gouse.ReverseDate())
	println("ShortDash:", gouse.DashDate())
	println("ShortDot:", gouse.DotDate())
	println("ShortUnderscore:", gouse.UnderlineDate())
	println("ShortSpace:", gouse.SpaceDate())
	println("ShortMonth:", gouse.MonthDate())
}

func SampleDateLong() {
	println("Long:", gouse.LongDate())
}

func SampleDateUTC() {
	println("UTC:", gouse.UTCDate())
}

func SampleDateToSecond() {
	println("ToSecond:", gouse.ToSecond(1))
}

func SampleDateToMinute() {
	println("ToMinute:", gouse.ToMinute(1))
}

func SampleDateToHour() {
	println("ToHour:", gouse.ToHour(1))
}

func SampleDateSleepSecond() {
	gouse.SleepSecond(1)
}

func SampleDateSleepMinute() {
	gouse.SleepMinute(1)
}

func SampleDateSleepHour() {
	gouse.SleepHour(1)
}

func SampleDateClock() {
	gouse.TerminalClock()
}
