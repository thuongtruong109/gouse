
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.25rem 0.75rem 0.25rem 0;' type='info' text='🔖 Date' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

### <Badge style='font-size: 1.1rem;' type='tip' text='1. sample date time' />



```go
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
```

### <Badge style='font-size: 1.1rem;' type='tip' text='2. sample date i s o' />



```go
func SampleDateISO() {
	println("ISO:", gouse.ISODate())
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='3. sample date short' />



```go
func SampleDateShort() {
	println("ShortNormal:", gouse.NormalDate())
	println("ShortReverse:", gouse.ReverseDate())
	println("ShortDash:", gouse.DashDate())
	println("ShortDot:", gouse.DotDate())
	println("ShortUnderscore:", gouse.UnderlineDate())
	println("ShortSpace:", gouse.SpaceDate())
	println("ShortMonth:", gouse.MonthDate())
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='4. sample date long' />



```go
func SampleDateLong() {
	println("Long:", gouse.LongDate())
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='5. sample date u t c' />



```go
func SampleDateUTC() {
	println("UTC:", gouse.UTCDate())
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='6. sample date to second' />



```go
func SampleDateToSecond() {
	println("ToSecond:", gouse.ToSecond(1))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='7. sample date to minute' />



```go
func SampleDateToMinute() {
	println("ToMinute:", gouse.ToMinute(1))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='8. sample date to hour' />



```go
func SampleDateToHour() {
	println("ToHour:", gouse.ToHour(1))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='9. sample date sleep second' />



```go
func SampleDateSleepSecond() {
	gouse.SleepSecond(1)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='10. sample date sleep minute' />



```go
func SampleDateSleepMinute() {
	gouse.SleepMinute(1)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='11. sample date sleep hour' />



```go
func SampleDateSleepHour() {
	gouse.SleepHour(1)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='12. sample date clock' />



```go
func SampleDateClock() {
	gouse.TerminalClock()
}
```
