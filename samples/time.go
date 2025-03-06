package samples

import "github.com/thuongtruong109/gouse"

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
	println("UnixNano:", gouse.UnixNano())
	println("UnixMilli:", gouse.UnixMilli())
	println("UnixMicro:", gouse.UnixMicro())
	gouse.Println("UnixMilliToTime:", gouse.UnixMilliToTime(1000000000))
	gouse.Println("UnixMicroToTime:", gouse.UnixMicroToTime(1000000000))
	gouse.Println("UnixNanoToTime:", gouse.UnixNanoToTime(1000000000))
}

/*
Description: Convert time to seconds
Input params: (time int64)
*/
func ToSecond() {
	println("ToSecond:", gouse.ToSecond(1))
}

/*
Description: Convert time to minutes
Input params: (time int64)
*/
func ToMinute() {
	println("ToMinute:", gouse.ToMinute(1))
}

/*
Description: Convert time to hours
Input params: (time int64)
*/
func ToHour() {
	println("ToHour:", gouse.ToHour(1))
}

/*
Description: Sleep for seconds
Input params: (interval int64)
*/
func SleepSecond() {
	gouse.SleepSecond(1)
}

/*
Description: Sleep for minutes
Input params: (interval int64)
*/
func SleepMinute() {
	gouse.SleepMinute(1)
}

/*
Description: Sleep for hours
Input params: (interval int64)
*/
func SleepHour() {
	gouse.SleepHour(1)
}

/*
Description: Display clock in terminal
*/
func Clock() {
	gouse.TerminalClock()
}
