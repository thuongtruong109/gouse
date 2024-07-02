# Read_by_line

## Imports

```go
import (
	"github.com/thuongtruong109/gouse/io"
)
```
## Functions


### SampleIoReadFileByLine

```go
func SampleIoReadFileByLine() {
	data, err := io.ReadFileByLine("main.go")
	if err != nil {
		println(err.Error())
	}
	for _, v := range data {
		println(v)
	}
}
```
