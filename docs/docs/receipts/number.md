
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Number' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Number clamp



```go
func NumberClamp() {
	println("Clamp number: ", gouse.Clamp(5, 1, 10))
}
```

## 2. Number check in range



```go
func NumberCheckInRange() {
	println("Check number is in range: ", gouse.InRange(5, 1, 10))
}
```
