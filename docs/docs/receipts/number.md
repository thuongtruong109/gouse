# Number

## Imports

```go
import (
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleNumClamp

```go
func SampleNumClamp() {
	println("Clamp number: ", gouse.Clamp(5, 1, 10))
}
```
## Imports

```go
import (
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleNumRandom

```go
func SampleNumRandom() {
	random := gouse.RandNum(1, 10)
	println("Random number [1, 10): ", random)
}
```
## Imports

```go
import (
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleNumInRange

```go
func SampleNumInRange() {
	println("Check number is in range: ", gouse.InRange(5, 1, 10))
}
```
