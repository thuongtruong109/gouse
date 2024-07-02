# Is_exist

## Imports

```go
import (
	"github.com/thuongtruong109/gouse/io"
)
```
## Functions


### SampleIoCheckFile

```go
func SampleIoCheckFile() {
	isExist, err := io.IsExistFile("data.json")
	if err != nil {
		println(err.Error())
	}
	if isExist {
		println("file exist")
	} else {
		println("file not exist")
	}
}
```
