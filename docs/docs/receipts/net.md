
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Net' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Net check u r l

Description: Check if a URL status<br>Input params: (url string)<br>

```go
func NetCheckURL() {
	isOk, statusCode, err := gouse.CheckUrl("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Response: ", isOk)
	println("Status code: ", statusCode)
}
```

## 2. Net connect time

Description: Calculate the time it takes to connect to a URL<br>Input params: (url string)<br>

```go
func NetConnectTime() {
	connectTime, err := gouse.ConTimeUrl("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connect time: %fs\n", connectTime)
}
```

## 3. Net encode u r l

Description: Encode a URL<br>Input params: (url string)<br>

```go
func NetEncodeURL() {
	println("Encode: ", gouse.EncodeUrl("https://google.com"))
}
```

## 4. Net decode u r l

Description: Decode a URL<br>Input params: (url string)<br>

```go
func NetDecodeURL() {
	println("Decode: ", gouse.DecodeUrl("https%3A%2F%2Fgoogle.com"))
}
```

## 5. Net header u r l

Description: Get the header of a URL<br>Input params: (url string)<br>

```go
func NetHeaderURL() {
	header, err := gouse.HeaderUrl("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(gouse.Map2Str(header))
}
```

## 6. Net port checker

Description: Check if a port is open<br>Input params: (protocol, hostname string, port int)<br>

```go
func NetPortChecker() {
	open := gouse.CheckPort("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
}
```

## 7. Net port scanner

Description: Scan for open ports on a given host.<br>Input params: (protocol, hostname string, startPort, endPort int)<br>

```go
func NetPortScanner() {
	gouse.ScanPort("tcp", "127.0.0.1", 3000, 8080)
}
```

## 8. Net proxy

Description: Proxy wrapper to another port<br>Input params: (port string, urls []string{})<br>

```go
func NetProxy() {
	gouse.Proxy("5000", []string{"http://localhost:3000", "http://localhost:3001"})
}
```

## 9. Net open u r l

Description: Open a URL in the default browser<br>Input params: (url string)<br>

```go
func NetOpenURL() {
	gouse.OpenUrl("https://google.com")
}
```
