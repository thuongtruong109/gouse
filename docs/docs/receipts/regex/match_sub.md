# Match_sub

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse/regex"
)
```
## Functions


### SampleRegexMatch

```go
func SampleRegexMatch() {
	fmt.Println("Match string with regex: ", regex.Match(`[A-Z]`, "Hello World 123"))
}
```