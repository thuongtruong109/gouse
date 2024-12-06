
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Helper' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Helper random i d



```go
func HelperRandomID() {
	println("Generate random ID: ", gouse.RandID())
}
```

## 2. Helper u u i d



```go
func HelperUUID() {
	println("New uuid: ", gouse.UUID())
}
```

## 3. Helper auto md doc



```go
func HelperAutoMdDoc() {
	inputFilePath := "main.go"
	outputFilePath := "main.md"
	gouse.Go2Md(inputFilePath, outputFilePath)
}
```
