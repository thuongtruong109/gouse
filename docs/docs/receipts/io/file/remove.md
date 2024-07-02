# Remove

## Imports

```go
import (
	"github.com/thuongtruong109/gouse/io"
)
```
## Functions


### SampleIoRemoveFile

```go
func SampleIoRemoveFile() {
	err := io.RemoveFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file removed")
}
```
