# Truncate

## Imports

```go
import (
	"github.com/thuongtruong109/gouse/io"
)
```
## Functions


### SampleIoTruncateFile

```go
func SampleIoTruncateFile() {
	err := io.TruncateFile("data.json", 10)
	if err != nil {
		println(err.Error())
	}
	println("file truncated to 10 bytes")
}
```
