
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.25rem 0.75rem 0.25rem 0;' type='info' text='🔖 Number' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

### <Badge style='font-size: 1.1rem;' type='tip' text='1. sample num clamp' />



```go
func SampleNumClamp() {
	println("Clamp number: ", gouse.Clamp(5, 1, 10))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='2. sample num random' />



```go
func SampleNumRandom() {
	random := gouse.RandNum(1, 10)
	println("Random number [1, 10): ", random)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='3. sample num in range' />



```go
func SampleNumInRange() {
	println("Check number is in range: ", gouse.InRange(5, 1, 10))
}
```
