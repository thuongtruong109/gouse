# Delay

## Imports

```go
import (
	"github.com/thuongtruong109/gouse/function"
)
```
## Functions


### SampleFuncDelay

```go
func SampleFuncDelay() {
	println("Delay start:")

	result := function.DelayF(func() string {
		return "Delayed return after 2s"
	}, 2)

	if result.HasReturn {
		println(result.Value)
	} else {
		println("No result")
	}

	function.Delay(func() {
		println("Delayed not return")
	}, 3)
}
```