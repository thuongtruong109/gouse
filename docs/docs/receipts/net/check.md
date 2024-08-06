# Check

## Imports

```go
import (
	"github.com/thuongtruong109/gouse/net")
```
## Functions


### SampleNetCheck

```go
func SampleNetCheck() {
	ok, err := net.Check("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Response: ", ok)
}```
## Imports

```go
import (
	"github.com/thuongtruong109/gouse/net")
```
## Functions


### SampleNetCheckWithStatusCode

```go
func SampleNetCheckWithStatusCode() {
	statusCode, err := net.CheckWithStatusCode("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Status code: ", statusCode)
}```
