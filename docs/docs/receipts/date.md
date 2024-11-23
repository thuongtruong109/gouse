# Date

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

#### 1. SampleDateTime

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

#### 2. SampleDateISO

```go
func SampleDateISO() {
	println("ISO:", gouse.ISODate())
}
```

#### 3. SampleDateShort

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

#### 4. SampleDateLong

```go
func SampleDateLong() {
	println("Long:", gouse.LongDate())
}
```

#### 5. SampleDateUTC

```go
func SampleDateUTC() {
	println("UTC:", gouse.UTCDate())
}
```

#### 6. SampleDateToSecond

```go
func SampleDateToSecond() {
	println("ToSecond:", gouse.ToSecond(1))
}
```

#### 7. SampleDateToMinute

```go
func SampleDateToMinute() {
	println("ToMinute:", gouse.ToMinute(1))
}
```

#### 8. SampleDateToHour

```go
func SampleDateToHour() {
	println("ToHour:", gouse.ToHour(1))
}
```

#### 9. SampleDateSleepSecond

```go
func SampleDateSleepSecond() {
	gouse.SleepSecond(1)
}
```

#### 10. SampleDateSleepMinute

```go
func SampleDateSleepMinute() {
	gouse.SleepMinute(1)
}
```

#### 11. SampleDateSleepHour

```go
func SampleDateSleepHour() {
	gouse.SleepHour(1)
}
```

#### 12. SampleDateClock

```go
func SampleDateClock() {
	gouse.TerminalClock()
}
```
