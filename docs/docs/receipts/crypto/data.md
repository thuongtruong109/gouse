# Data

## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/crypto")
```
## Functions


### SampleCryptoEncode

```go
func SampleCryptoEncode() {
	data := []byte("This is a sample data")

	encodedData, err := crypto.Encode(data)
	if err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}
	fmt.Println("Raw data:", string(data))
	fmt.Println("Encoded data:", string(encodedData))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/crypto")
```
## Functions


### SampleCryptoDecode

```go
func SampleCryptoDecode() {
	data := []byte("VGhpcyBpcyBhIHNhbXBsZSBkYXRh")
	decodedData, err := crypto.Decode(data)
	if err != nil {
		fmt.Println("Error decoding data:", err)
		return
	}
	fmt.Println("Raw encoded data:", string(data))
	fmt.Println("Decoded data:", string(decodedData))
}```
