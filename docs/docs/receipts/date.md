
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Date' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Date time



```go
func DateTime() {
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
```

## 2. Date i s o



```go
func DateISO() {
	println("ISO:", gouse.ISODate())
}
```

## 3. Date short



```go
func DateShort() {
	println("ShortNormal:", gouse.NormalDate())
	println("ShortReverse:", gouse.ReverseDate())
	println("ShortDash:", gouse.DashDate())
	println("ShortDot:", gouse.DotDate())
	println("ShortUnderscore:", gouse.UnderlineDate())
	println("ShortSpace:", gouse.SpaceDate())
	println("ShortMonth:", gouse.MonthDate())
}
```

## 4. Date long



```go
func DateLong() {
	println("Long:", gouse.LongDate())
}
```

## 5. Date u t c



```go
func DateUTC() {
	println("UTC:", gouse.UTCDate())
}
```

## 6. Date to second



```go
func DateToSecond() {
	println("ToSecond:", gouse.ToSecond(1))
}
```

## 7. Date to minute



```go
func DateToMinute() {
	println("ToMinute:", gouse.ToMinute(1))
}
```

## 8. Date to hour



```go
func DateToHour() {
	println("ToHour:", gouse.ToHour(1))
}
```

## 9. Date sleep second



```go
func DateSleepSecond() {
	gouse.SleepSecond(1)
}
```

## 10. Date sleep minute



```go
func DateSleepMinute() {
	gouse.SleepMinute(1)
}
```

## 11. Date sleep hour



```go
func DateSleepHour() {
	gouse.SleepHour(1)
}
```

## 12. Date clock



```go
func DateClock() {
	gouse.TerminalClock()
}
```
