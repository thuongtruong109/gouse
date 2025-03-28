package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Show all time elements
*/
func TimeElement() {
	println("Second:", gouse.Second())
	println("Minute:", gouse.Minute())
	println("Hour:", gouse.Hour())
	println("Day:", gouse.Day())
	println("Month:", gouse.Month())
	println("Year:", gouse.Year())
	println("Weekday:", gouse.Weekday())
	println("Unix:", gouse.Unix())
	println("UnixNano:", gouse.Nano())
	println("UnixMilli:", gouse.Milli())
	println("UnixMicro:", gouse.Micro())
	fmt.Println("UnixMilliToTime:", gouse.Milli2Time(1000000000))
	fmt.Println("UnixMicroToTime:", gouse.Micro2Time(1000000000))
	fmt.Println("UnixNanoToTime:", gouse.Nano2Time(1000000000))
}

/*
Description: Convert time to seconds
Input params: (time int64)
*/
func TimeToSecond() {
	println("ToSecond:", gouse.ToSecond(1))
}

/*
Description: Convert time to minutes
Input params: (time int64)
*/
func TimeToMinute() {
	println("ToMinute:", gouse.ToMinute(1))
}

/*
Description: Convert time to hours
Input params: (time int64)
*/
func TimeToHour() {
	println("ToHour:", gouse.ToHour(1))
}

/*
Description: Sleep for seconds
Input params: (interval int64)
*/
func TimeSleepSecond() {
	gouse.SleepS(1)
}

/*
Description: Sleep for minutes
Input params: (interval int64)
*/
func TimeSleepMinute() {
	gouse.SleepM(1)
}

/*
Description: Sleep for hours
Input params: (interval int64)
*/
func TimeSleepHour() {
	gouse.SleepH(1)
}

/*
Description: Display clock in terminal
*/
func TimeTerminalClock() {
	gouse.TerminalClock()
}
