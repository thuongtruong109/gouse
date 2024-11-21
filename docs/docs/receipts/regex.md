# Regex

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleRegexIsMatch

```go
func SampleRegexIsMatch() {
	fmt.Println("Match string with regex: ", gouse.IsMatchReg(`[a-z]+\s[a-z]+`, "hello world"))
}
```
## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleRegexMatchIndex

```go
func SampleRegexMatchIndex() {
	paragraph := "I think Ruth's dog is cuter than your dog!"
	matchIdx := gouse.MatchIndexReg(`[^\w\s']`, paragraph)
	if matchIdx != -1 {
		fmt.Printf("Match with regex (index: %d, value: %s)\n", matchIdx, string(paragraph[matchIdx]))
	} else {
		println("Not found index match regex")
	}
}
```
## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleRegexMatch

```go
func SampleRegexMatch() {
	fmt.Println("Match string with regex: ", gouse.MatchReg(`[A-Z]`, "Hello World 123"))
}
```
