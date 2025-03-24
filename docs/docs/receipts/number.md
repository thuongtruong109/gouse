
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Number' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Number clamp

Description: Clamp a number between a minimum and maximum value<br>Input params: (number, min, max int) -> int<br>

```go
func NumberClamp() {
	println("Clamp number: ", gouse.Clamp(5, 1, 10))
}
```

## 2. Number in range

Description: Check if a number is in a range<br>Input params: (number, min, max int) -> bool<br>

```go
func NumberInRange() {
	println("Check number is in range: ", gouse.InRange(5, 1, 10))
}
```

## 3. Number sort

Description: Sort a list of numbers<br>Input params: (nums []int) -> []int<br>

```go
func NumberSort() {
	nums := []int{5, 3, 8, 1, 2}
	println("Sort numbers: ", gouse.SortNum(nums))
}
```
