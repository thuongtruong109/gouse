# Number

```go
import (
	"github.com/thuongtruong109/gouse"
)
```

#### 1. SampleNumClamp

```go
func SampleNumClamp() {
	println("Clamp number: ", gouse.Clamp(5, 1, 10))
}
```

#### 2. SampleNumRandom

```go
func SampleNumRandom() {
	random := gouse.RandNum(1, 10)
	println("Random number [1, 10): ", random)
}
```

#### 3. SampleNumInRange

```go
func SampleNumInRange() {
	println("Check number is in range: ", gouse.InRange(5, 1, 10))
}
```
