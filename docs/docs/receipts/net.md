# Net

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

#### 1. SampleNetCheck

```go
func SampleNetCheck() {
	ok, err := gouse.CheckHref("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Response: ", ok)
}
```

#### 2. SampleNetCheckWithStatusCode

```go
func SampleNetCheckWithStatusCode() {
	statusCode, err := gouse.CheckHrefStatusCode("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Status code: ", statusCode)
}
```

#### 3. SampleNetConnectTime

```go
func SampleNetConnectTime() {
	connectTime, err := gouse.HrefConnectTime("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connect time: %fs\n", connectTime)
}
```

#### 4. SampleNetEncode

```go
func SampleNetEncode() {
	println("Encode: ", gouse.EncodeHref("https://google.com"))
}
```

#### 5. SampleNetDecode

```go
func SampleNetDecode() {
	println("Decode: ", gouse.DecodeHref("https%3A%2F%2Fgoogle.com"))
}
```

#### 6. SampleNetHeader

```go
func SampleNetHeader() {
	header, err := gouse.HrefHeader("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(gouse.MapAsString(header))
}
```

#### 7. SampleApiPortChecker

```go
func SampleApiPortChecker() {
	open := gouse.CheckPort("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
}
```

#### 8. SampleApiPortScanner

```go
func SampleApiPortScanner() {
	gouse.ScanPort("tcp", "127.0.0.1", 3000, 8080)
}
```

#### 9. SampleNetProxy

```go
func SampleNetProxy() {
	gouse.Proxy("5000", []string{"http://localhost:3000", "http://localhost:3001"})
}
```

#### 10. SampleNetOpen

```go
func SampleNetOpen() {
	gouse.OpenHref("https://google.com")
}
```
