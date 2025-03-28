
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Utilities' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Utils go to markdown

Description: Auto generate markdown document from go source code<br>Input params: (inputFilePath: string, outputFilePath: string)<br>

```go
func UtilsGoToMarkdown() {
	inputFilePath := "main.go"
	outputFilePath := "main.md"
	gouse.Go2Md(inputFilePath, outputFilePath)
}
```
