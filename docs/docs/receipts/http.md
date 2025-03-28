
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Http' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Http check url

Description: Check if a URL status<br>Input params: (url string)<br>

```go
func HttpCheckUrl() {
	isOk, statusCode, err := gouse.CheckUrl("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Response: ", isOk)
	println("Status code: ", statusCode)
}
```

## 2. Http connect time

Description: Calculate the time it takes to connect to a URL<br>Input params: (url string)<br>

```go
func HttpConnectTime() {
	connectTime, err := gouse.ConTimeUrl("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connect time: %fs\n", connectTime)
}
```

## 3. Http encode url

Description: Encode a URL<br>Input params: (url string)<br>

```go
func HttpEncodeUrl() {
	println("Encode: ", gouse.EncodeUrl("https://google.com"))
}
```

## 4. Http decode url

Description: Decode a URL<br>Input params: (url string)<br>

```go
func HttpDecodeUrl() {
	println("Decode: ", gouse.DecodeUrl("https%3A%2F%2Fgoogle.com"))
}
```

## 5. Http header url

Description: Get the header of a URL<br>Input params: (url string)<br>

```go
func HttpHeaderUrl() {
	header, err := gouse.HeaderUrl("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(gouse.Map2Str(header))
}
```

## 6. Http port checker

Description: Check if a port is open<br>Input params: (protocol, hostname string, port int)<br>

```go
func HttpPortChecker() {
	open := gouse.CheckPort("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
}
```

## 7. Http port scanner

Description: Scan for open ports on a given host.<br>Input params: (protocol, hostname string, startPort, endPort int)<br>

```go
func HttpPortScanner() {
	gouse.ScanPort("tcp", "127.0.0.1", 3000, 8080)
}
```

## 8. Http open url

Description: Open a URL in the default browser<br>Input params: (url string)<br>

```go
func HttpOpenUrl() {
	gouse.OpenUrl("https://google.com")
}
```
