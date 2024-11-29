
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.25rem 0.75rem 0.25rem 0;' type='info' text='🔖 Helper' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

### <Badge style='font-size: 1.1rem;' type='tip' text='1. sample helper random i d' />



```go
func SampleHelperRandomID() {
	println("Generate random ID: ", gouse.RandID())
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='2. sample helper u u i d' />



```go
func SampleHelperUUID() {
	println("New uuid: ", gouse.UUID())
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='3. sample helper auto md doc' />



```go
func SampleHelperAutoMdDoc() {
	inputFilePath := "main.go"
	outputFilePath := "main.md"
	gouse.Go2Md(inputFilePath, outputFilePath)
}
```
