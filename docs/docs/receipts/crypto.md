# Crypto

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```
## Functions


### SampleCryptoEncode

```go
func SampleCryptoEncode() {
	data := []byte("This is a sample data")

	encodedData, err := gouse.EncodeData(data)
	if err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}
	fmt.Println("Raw data:", string(data))
	fmt.Println("Encoded data:", string(encodedData))
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


### SampleCryptoDecode

```go
func SampleCryptoDecode() {
	data := []byte("VGhpcyBpcyBhIHNhbXBsZSBkYXRh")
	decodedData, err := gouse.DecodeData(data)
	if err != nil {
		fmt.Println("Error decoding data:", err)
		return
	}
	fmt.Println("Raw encoded data:", string(data))
	fmt.Println("Decoded data:", string(decodedData))
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


### SampleCryptoEncryptFile

```go
func SampleCryptoEncryptFile() {
	gouse.EncryptFile("sample.txt", []byte("password"))
	println("File content encrypted")
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


### SampleCryptoDecryptFile

```go
func SampleCryptoDecryptFile() {
	gouse.DecryptFile("sample.txt", []byte("password"))
	println("File content decrypted")
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


### SampleCryptoEncryptPassword

```go
func SampleCryptoEncryptPassword() {
	data := "This is a sample data"

	encryptedData, err := gouse.EncryptPassword(data)
	if err != nil {
		fmt.Println("Error encrypting data:", err)
		return
	}
	fmt.Println("Raw data:", string(data))
	fmt.Println("Encrypted data:", string(encryptedData))
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


### SampleCryptoDecryptPassword

```go
func SampleCryptoDecryptPassword() {
	data := "$2a$10$bcA002IOHi5SYHNH4lmIbuHjHplGl7TQZ.MznNrL1N70vAi7ovTa2"
	err := gouse.DecryptPassword(data, "This is a sample data")
	if err != nil {
		fmt.Println("Error decrypting data:", err)
		return
	}
	println("Password matched")
}
```
