
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Time' />


```go
import (
	"fmt"
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
	println("UnixNano:", gouse.Nano())
	println("UnixMilli:", gouse.Milli())
	println("UnixMicro:", gouse.Micro())
	fmt.Println("UnixMilliToTime:", gouse.Milli2Time(1000000000))
	fmt.Println("UnixMicroToTime:", gouse.Micro2Time(1000000000))
	fmt.Println("UnixNanoToTime:", gouse.Nano2Time(1000000000))
}
```

## 2. Time to second

Description: Convert time to seconds<br>Input params: (time int64)<br>

```go
func TimeToSecond() {
	println("ToSecond:", gouse.ToSecond(1))
}
```

## 3. Time to minute

Description: Convert time to minutes<br>Input params: (time int64)<br>

```go
func TimeToMinute() {
	println("ToMinute:", gouse.ToMinute(1))
}
```

## 4. Time to hour

Description: Convert time to hours<br>Input params: (time int64)<br>

```go
func TimeToHour() {
	println("ToHour:", gouse.ToHour(1))
}
```

## 5. Time sleep second

Description: Sleep for seconds<br>Input params: (interval int64)<br>

```go
func TimeSleepSecond() {
	gouse.SleepS(1)
}
```

## 6. Time sleep minute

Description: Sleep for minutes<br>Input params: (interval int64)<br>

```go
func TimeSleepMinute() {
	gouse.SleepM(1)
}
```

## 7. Time sleep hour

Description: Sleep for hours<br>Input params: (interval int64)<br>

```go
func TimeSleepHour() {
	gouse.SleepH(1)
}
```

## 8. Time terminal clock

Description: Display clock in terminal<br>

```go
func TimeTerminalClock() {
	gouse.TerminalClock()
}
```
