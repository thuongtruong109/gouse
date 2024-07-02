# Retry

## Imports

```go
import (
	"github.com/thuongtruong109/gouse/function"
)
```
## Functions


### SampleFuncRetry

```go
func SampleFuncRetry() {
	function.Retry(func() error {
		println("Retry")
		return nil
	}, 3, 1)
}
```
