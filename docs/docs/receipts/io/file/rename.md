# Rename

## Imports

```go
import (
	"github.com/thuongtruong109/gouse/io")
```
## Functions


### SampleIoRenameFile

```go
func SampleIoRenameFile() {
	err := io.RenameFile("data.json", "data2.json")
	if err != nil {
		println(err.Error())
	}
	println("file renamed")
}```
