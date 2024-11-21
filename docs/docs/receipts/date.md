# Date

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleDateTime

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
## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleDateISO

```go
func SampleDateISO() {
	println("ISO:", gouse.ISODate())
}
```
## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleDateShort

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
## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleDateLong

```go
func SampleDateLong() {
	println("Long:", gouse.LongDate())
}
```
## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleDateUTC

```go
func SampleDateUTC() {
	println("UTC:", gouse.UTCDate())
}
```
## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleDateToSecond

```go
func SampleDateToSecond() {
	println("ToSecond:", gouse.ToSecond(1))
}
```
## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleDateToMinute

```go
func SampleDateToMinute() {
	println("ToMinute:", gouse.ToMinute(1))
}
```
## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleDateToHour

```go
func SampleDateToHour() {
	println("ToHour:", gouse.ToHour(1))
}
```
## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleDateSleepSecond

```go
func SampleDateSleepSecond() {
	gouse.SleepSecond(1)
}
```
## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleDateSleepMinute

```go
func SampleDateSleepMinute() {
	gouse.SleepMinute(1)
}
```
## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleDateSleepHour

```go
func SampleDateSleepHour() {
	gouse.SleepHour(1)
}
```
## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleDateClock

```go
func SampleDateClock() {
	gouse.TerminalClock()
}
```
