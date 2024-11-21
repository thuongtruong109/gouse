# Helper

## Imports

```go
import (
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleHelperRandomID

```go
func SampleHelperRandomID() {
	println("Generate random ID: ", gouse.RandID())
}
```
## Imports

```go
import (
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleHelperUUID

```go
func SampleHelperUUID() {
	println("New uuid: ", gouse.UUID())
}
```
## Imports

```go
import (
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleHelperAutoMdDoc

```go
func SampleHelperAutoMdDoc() {
	inputFilePath := "main.go"
	outputFilePath := "main.md"
	gouse.AutoMdDoc(inputFilePath, outputFilePath)
}
```
