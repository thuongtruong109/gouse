# Root

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse/math"
)
```
## Functions


### SampleMathRoot

```go
func SampleMathRoot() {
	println("Square-Root of integer number: ", math.Sqrt(16))
	println("Square-Root of float number: ", fmt.Sprintf("%f", math.SqrtF(20.0)))

	println("Cube-Root of integer number: ", math.Cbrt(27))
	println("Cube-Root of float number: ", fmt.Sprintf("%f", math.CbrtF(20.0)))

	println("Nth-Root of integer number: ", math.Root(4, 2))
	println("Nth-Root of float number: ", fmt.Sprintf("%f", math.RootF(20.0, 3.0)))
}
```