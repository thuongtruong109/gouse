# Circle

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse/math"
)
```
## Functions


### SampleMathCircle

```go
func SampleMathCircle() {
	println("Area of circle (integer): ", fmt.Sprintf("%f", math.AreaCircle(10)))
	println("Area of circle (float): ", fmt.Sprintf("%f", math.AreaCircleF(10.0)))
	println("Perimeter of circle (integer): ", fmt.Sprintf("%f", math.PeriCircle(10)))
	println("Perimeter of circle (float): ", fmt.Sprintf("%f", math.PeriCircleF(10.0)))
}
```