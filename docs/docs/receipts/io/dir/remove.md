# Remove

## Imports

```go
import (
	"github.com/thuongtruong109/gouse/io")
```
## Functions


### SampleIoRemoveDir

```go
func SampleIoRemoveDir() {
	err3 := io.RemoveDir("tmp")
	if err3 != nil {
		println(err3.Error())
	}
	println("dir removed")
}```
