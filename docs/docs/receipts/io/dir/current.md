# Current

## Imports

```go
import (
	"github.com/thuongtruong109/gouse/io")
```
## Functions


### SampleIoCurrentDir

```go
func SampleIoCurrentDir() {
	data, err := io.CurrentDir()
	if err != nil {
		println(err.Error())
		return
	}

	println(data)
}```
