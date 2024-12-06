
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Crypto' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Crypto encode

Description: Encode data to base64<br>Input params: (data []byte)<br>

```go
func CryptoEncode() {
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

## 2. Crypto decode

Description: Decode data from base64<br>Input params: (data []byte)<br>

```go
func CryptoDecode() {
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

## 3. Crypto encrypt file

Description: Encrypt data in file<br>Input params: (filename string, password []byte)<br>

```go
func CryptoEncryptFile() {
	gouse.EncryptFile("sample.txt", []byte("password"))
	println("File content encrypted")
}
```

## 4. Crypto decrypt file

Description: Decrypt data in file<br>Input params: (filename string, password []byte)<br>

```go
func CryptoDecryptFile() {
	gouse.DecryptFile("sample.txt", []byte("password"))
	println("File content decrypted")
}
```

## 5. Crypto encrypt password

Description: Encrypt password string<br>Input params: (data string)<br>

```go
func CryptoEncryptPassword() {
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

## 6. Crypto decrypt password

Description: Decrypt password string and compare with the original password<br>Input params: (data string, password string)<br>

```go
func CryptoDecryptPassword() {
	data := "$2a$10$bcA002IOHi5SYHNH4lmIbuHjHplGl7TQZ.MznNrL1N70vAi7ovTa2"
	err := gouse.DecryptPassword(data, "This is a sample data")
	if err != nil {
		fmt.Println("Error decrypting data:", err)
		return
	}
	println("Password matched")
}
```
