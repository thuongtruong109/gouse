# Create

## Imports

```go
import (
	"github.com/thuongtruong109/gouse/io")
```
## Functions


### SampleIoCreateFile

```go
func SampleIoCreateFile() {
	err := io.CreateFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file created")
}```
