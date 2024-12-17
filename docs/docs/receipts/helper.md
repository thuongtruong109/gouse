
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Helper' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. U u i d

Description: Generate a new UUID<br>

```go
func UUID() {
	println("New uuid: ", gouse.UUID())
}
```

## 2. Auto markdown document

Description: Auto generate markdown document from go source code<br>Input params: (inputFilePath: string, outputFilePath: string)<br>

```go
func AutoMarkdownDocument() {
	inputFilePath := "main.go"
	outputFilePath := "main.md"
	gouse.Go2Md(inputFilePath, outputFilePath)
}
```
