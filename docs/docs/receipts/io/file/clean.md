# Clean

## Imports

```go
import (
	"github.com/thuongtruong109/gouse/io"
)
```
## Functions


### SampleIoCleanFile

```go
func SampleIoCleanFile() {
	err := io.CleanFile("data.json")
	if err != nil {
		println(err.Error())
	}

	// or using truncate with 0 bytes
	// err := helper.TruncateFile("data.json", 0)
	// if err != nil {
	// 	println(err.Error())
	// }

	println("file cleaned")
}
```
