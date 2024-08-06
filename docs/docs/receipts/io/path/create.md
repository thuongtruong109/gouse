# Create

## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/io")
```
## Functions


### SampleIoCreatePath

```go
func SampleIoCreatePath() {
	relativePath := "tmp/example.txt"

	if err := io.CreatePath(relativePath); err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	println("File created successfully.")
}```
