
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Date' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

### <Badge style='font-size: 1.1rem;' type='tip' text='1. date time' />



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

### <Badge style='font-size: 1.1rem;' type='tip' text='2. date i s o' />



```go
func DateISO() {
	println("ISO:", gouse.ISODate())
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='3. date short' />



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

### <Badge style='font-size: 1.1rem;' type='tip' text='4. date long' />



```go
func DateLong() {
	println("Long:", gouse.LongDate())
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='5. date u t c' />



```go
func DateUTC() {
	println("UTC:", gouse.UTCDate())
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='6. date to second' />



```go
func DateToSecond() {
	println("ToSecond:", gouse.ToSecond(1))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='7. date to minute' />



```go
func DateToMinute() {
	println("ToMinute:", gouse.ToMinute(1))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='8. date to hour' />



```go
func DateToHour() {
	println("ToHour:", gouse.ToHour(1))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='9. date sleep second' />



```go
func DateSleepSecond() {
	gouse.SleepSecond(1)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='10. date sleep minute' />



```go
func DateSleepMinute() {
	gouse.SleepMinute(1)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='11. date sleep hour' />



```go
func DateSleepHour() {
	gouse.SleepHour(1)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='12. date clock' />



```go
func DateClock() {
	gouse.TerminalClock()
}
```
