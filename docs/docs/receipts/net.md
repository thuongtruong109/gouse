
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.25rem 0.75rem 0.25rem 0;' type='info' text='🔖 Net' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

### <Badge style='font-size: 1.1rem;' type='tip' text='1. sample net check' />



```go
func SampleNetCheck() {
	ok, err := gouse.CheckHref("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Response: ", ok)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='2. sample net check with status code' />

Description: Check if a URL is valid<br>Input params: (url)<br>

```go
func SampleNetCheckWithStatusCode() {
	statusCode, err := gouse.CheckHrefStatusCode("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Status code: ", statusCode)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='3. sample net connect time' />

Description: Calculate the time it takes to connect to a URL<br>Input params: (url)<br>

```go
func SampleNetConnectTime() {
	connectTime, err := gouse.HrefConnectTime("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connect time: %fs\n", connectTime)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='4. sample net encode' />

Description: Encode a URL<br>Input params: (url)<br>

```go
func SampleNetEncode() {
	println("Encode: ", gouse.EncodeHref("https://google.com"))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='5. sample net decode' />

Description: Decode a URL<br>Input params: (url)<br>

```go
func SampleNetDecode() {
	println("Decode: ", gouse.DecodeHref("https%3A%2F%2Fgoogle.com"))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='6. sample net header' />

Description: Get the header of a URL<br>Input params: (url)<br>

```go
func SampleNetHeader() {
	header, err := gouse.HrefHeader("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(gouse.MapAsString(header))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='7. sample api port checker' />

Description: Check if a port is open<br>Input params: (protocol, hostname, port)<br>

```go
func SampleApiPortChecker() {
	open := gouse.CheckPort("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='8. sample api port scanner' />

Description: Scan for open ports on a given host.<br>Input params: (protocol, hostname, start port, end port)<br>

```go
func SampleApiPortScanner() {
	gouse.ScanPort("tcp", "127.0.0.1", 3000, 8080)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='9. sample net proxy' />

Description: Proxy wrapper to another port<br>Input params: (port, []string{urls})<br>

```go
func SampleNetProxy() {
	gouse.Proxy("5000", []string{"http://localhost:3000", "http://localhost:3001"})
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='10. sample net open' />

Description: Open a URL in the default browser<br>Input params: (url)<br>

```go
func SampleNetOpen() {
	gouse.OpenHref("https://google.com")
}
```
