
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Net' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Check u r l

Description: Check if a URL status<br>Input params: (url string)<br>

```go
func CheckURL() {
	isOk, statusCode, err := gouse.CheckUrl("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Response: ", isOk)
	println("Status code: ", statusCode)
}
```

## 2. Connect time

Description: Calculate the time it takes to connect to a URL<br>Input params: (url string)<br>

```go
func ConnectTime() {
	connectTime, err := gouse.ConTimeUrl("https://google.com")
	if err != nil {
		panic(err)
	}

	gouse.Printf("Connect time: %fs\n", connectTime)
}
```

## 3. Encode u r l

Description: Encode a URL<br>Input params: (url string)<br>

```go
func EncodeURL() {
	println("Encode: ", gouse.EncodeUrl("https://google.com"))
}
```

## 4. Decode u r l

Description: Decode a URL<br>Input params: (url string)<br>

```go
func DecodeURL() {
	println("Decode: ", gouse.DecodeUrl("https%3A%2F%2Fgoogle.com"))
}
```

## 5. Header u r l

Description: Get the header of a URL<br>Input params: (url string)<br>

```go
func HeaderURL() {
	header, err := gouse.HeaderUrl("https://google.com")
	if err != nil {
		panic(err)
	}

	gouse.Println(gouse.Map2Str(header))
}
```

## 6. Port checker

Description: Check if a port is open<br>Input params: (protocol, hostname string, port int)<br>

```go
func PortChecker() {
	open := gouse.CheckPort("tcp", "localhost", 1313)
	gouse.Printf("Port Open: %t\n", open)
}
```

## 7. Port scanner

Description: Scan for open ports on a given host.<br>Input params: (protocol, hostname string, startPort, endPort int)<br>

```go
func PortScanner() {
	gouse.ScanPort("tcp", "127.0.0.1", 3000, 8080)
}
```

## 8. Proxy

Description: Proxy wrapper to another port<br>Input params: (port string, urls []string{})<br>

```go
func Proxy() {
	gouse.Proxy("5000", []string{"http://localhost:3000", "http://localhost:3001"})
}
```

## 9. Open u r l

Description: Open a URL in the default browser<br>Input params: (url string)<br>

```go
func OpenURL() {
	gouse.OpenUrl("https://google.com")
}
```
