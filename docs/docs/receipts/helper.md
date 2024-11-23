# Helper

```go
import (
	"github.com/thuongtruong109/gouse"
)
```

#### 1. SampleHelperRandomID

```go
func SampleHelperRandomID() {
	println("Generate random ID: ", gouse.RandID())
}
```

#### 2. SampleHelperUUID

```go
func SampleHelperUUID() {
	println("New uuid: ", gouse.UUID())
}
```

#### 3. SampleHelperAutoMdDoc

```go
func SampleHelperAutoMdDoc() {
	inputFilePath := "main.go"
	outputFilePath := "main.md"
	gouse.Go2Md(inputFilePath, outputFilePath)
}
```
