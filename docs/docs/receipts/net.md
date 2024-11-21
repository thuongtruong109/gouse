# Net

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleNetCheck

```go
func SampleNetCheck() {
	ok, err := gouse.CheckHref("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Response: ", ok)
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


### SampleNetCheckWithStatusCode

```go
func SampleNetCheckWithStatusCode() {
	statusCode, err := gouse.CheckHrefStatusCode("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Status code: ", statusCode)
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


### SampleNetConnectTime

```go
func SampleNetConnectTime() {
	connectTime, err := gouse.HrefConnectTime("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connect time: %fs\n", connectTime)
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


### SampleNetEncode

```go
func SampleNetEncode() {
	println("Encode: ", gouse.EncodeHref("https://google.com"))
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


### SampleNetDecode

```go
func SampleNetDecode() {
	println("Decode: ", gouse.DecodeHref("https%3A%2F%2Fgoogle.com"))
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


### SampleNetHeader

```go
func SampleNetHeader() {
	header, err := gouse.HrefHeader("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(gouse.MapAsString(header))
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


### SampleApiPortChecker

```go
func SampleApiPortChecker() {
	open := gouse.CheckPort("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
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


### SampleApiPortScanner

```go
func SampleApiPortScanner() {
	gouse.ScanPort("tcp", "127.0.0.1", 3000, 8080)
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


### SampleNetProxy

```go
func SampleNetProxy() {
	gouse.Proxy("5000", []string{"http://localhost:3000", "http://localhost:3001"})
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


### SampleNetOpen

```go
func SampleNetOpen() {
	gouse.OpenHref("https://google.com")
}
```
