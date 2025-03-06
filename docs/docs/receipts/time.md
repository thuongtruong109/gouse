
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Time' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Time element

Description: Show all time elements<br>

```go
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
```

## 2. To second

Description: Convert time to seconds<br>Input params: (time int64)<br>

```go
func ToSecond() {
	println("ToSecond:", gouse.ToSecond(1))
}
```

## 3. To minute

Description: Convert time to minutes<br>Input params: (time int64)<br>

```go
func ToMinute() {
	println("ToMinute:", gouse.ToMinute(1))
}
```

## 4. To hour

Description: Convert time to hours<br>Input params: (time int64)<br>

```go
func ToHour() {
	println("ToHour:", gouse.ToHour(1))
}
```

## 5. Sleep second

Description: Sleep for seconds<br>Input params: (interval int64)<br>

```go
func SleepSecond() {
	gouse.SleepSecond(1)
}
```

## 6. Sleep minute

Description: Sleep for minutes<br>Input params: (interval int64)<br>

```go
func SleepMinute() {
	gouse.SleepMinute(1)
}
```

## 7. Sleep hour

Description: Sleep for hours<br>Input params: (interval int64)<br>

```go
func SleepHour() {
	gouse.SleepHour(1)
}
```

## 8. Clock

Description: Display clock in terminal<br>

```go
func Clock() {
	gouse.TerminalClock()
}
```
