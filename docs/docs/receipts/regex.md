
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Regex' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Regex is match

Description: Check if the input string is match with regex pattern<br>Input params: (pattern, input string)<br>--> Note: regex pattern is not include one of (^, $, /g)<br>

```go
func RegexIsMatch() {
	fmt.Println("Match string with regex: ", gouse.IsMatchReg(`[a-z]+\s[a-z]+`, "hello world"))
}
```

## 2. Regex match index

Description: Find the first index and value of the match regex pattern<br>Input params: (pattern, input string)<br>--> Note: regex pattern is not include one of (^, $, /g)<br>

```go
func RegexMatchIndex() {
	paragraph := "I think Ruth's dog is cuter than your dog!"
	matchIdx := gouse.MatchIdxReg(`[^\w\s']`, paragraph)
	if matchIdx != -1 {
		fmt.Printf("Match with regex (index: %d, value: %s)\n", matchIdx, string(paragraph[matchIdx]))
	} else {
		println("Not found index match regex")
	}
}
```

## 3. Regex match

Description: Find all index and value of the match regex pattern<br>Input params: (pattern, input string)<br>--> Note: regex pattern is not include one of (^, $, /g)<br>

```go
func RegexMatch() {
	fmt.Println("Match string with regex: ", gouse.MatchReg(`[A-Z]`, "Hello World 123"))
}
```
