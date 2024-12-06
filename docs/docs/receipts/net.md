
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Net' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Net check



```go
func NetCheck() {
	ok, err := gouse.CheckHref("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Response: ", ok)
}
```

## 2. Net check with status code

Description: Check if a URL is valid<br>Input params: (url)<br>

```go
func NetCheckWithStatusCode() {
	statusCode, err := gouse.CheckHrefStatusCode("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Status code: ", statusCode)
}
```

## 3. Net connect time

Description: Calculate the time it takes to connect to a URL<br>Input params: (url)<br>

```go
func NetConnectTime() {
	connectTime, err := gouse.HrefConnectTime("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connect time: %fs\n", connectTime)
}
```

## 4. Net encode

Description: Encode a URL<br>Input params: (url)<br>

```go
func NetEncode() {
	println("Encode: ", gouse.EncodeHref("https://google.com"))
}
```

## 5. Net decode

Description: Decode a URL<br>Input params: (url)<br>

```go
func NetDecode() {
	println("Decode: ", gouse.DecodeHref("https%3A%2F%2Fgoogle.com"))
}
```

## 6. Net header

Description: Get the header of a URL<br>Input params: (url)<br>

```go
func NetHeader() {
	header, err := gouse.HrefHeader("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(gouse.MapAsString(header))
}
```

## 7. Api port checker

Description: Check if a port is open<br>Input params: (protocol, hostname, port)<br>

```go
func ApiPortChecker() {
	open := gouse.CheckPort("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
}
```

## 8. Api port scanner

Description: Scan for open ports on a given host.<br>Input params: (protocol, hostname, start port, end port)<br>

```go
func ApiPortScanner() {
	gouse.ScanPort("tcp", "127.0.0.1", 3000, 8080)
}
```

## 9. Net proxy

Description: Proxy wrapper to another port<br>Input params: (port, []string{urls})<br>

```go
func NetProxy() {
	gouse.Proxy("5000", []string{"http://localhost:3000", "http://localhost:3001"})
}
```

## 10. Net open

Description: Open a URL in the default browser<br>Input params: (url)<br>

```go
func NetOpen() {
	gouse.OpenHref("https://google.com")
}
```
