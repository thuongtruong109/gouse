package samples

import "github.com/thuongtruong109/gouse"

/*
Description: Get current date in ISO format
*/
func DateISO() {
	println("ISO:", gouse.ISODate())
}

/*
Description: Get current date in short formats
*/
func DateShort() {
	println("ShortNormal:", gouse.NormalDate())
	println("ShortReverse:", gouse.ReverseDate())
	println("ShortDash:", gouse.DashDate())
	println("ShortDot:", gouse.DotDate())
	println("ShortUnderscore:", gouse.UnderlineDate())
	println("ShortSpace:", gouse.SpaceDate())
	println("ShortMonth:", gouse.MonthDate())
}

/*
Description: Get current date in long format
*/
func DateLong() {
	println("Long:", gouse.LongDate())
}

/*
Description: Get current date in UTC format
*/
func DateUTC() {
	println("UTC:", gouse.UTCDate())
}
