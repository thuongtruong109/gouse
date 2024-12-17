
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Date' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Date i s o

Description: Get current date in ISO format<br>

```go
func DateISO() {
	println("ISO:", gouse.ISODate())
}
```

## 2. Date short

Description: Get current date in short formats<br>

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

## 3. Date long

Description: Get current date in long format<br>

```go
func DateLong() {
	println("Long:", gouse.LongDate())
}
```

## 4. Date u t c

Description: Get current date in UTC format<br>

```go
func DateUTC() {
	println("UTC:", gouse.UTCDate())
}
```
