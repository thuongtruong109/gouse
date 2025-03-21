
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Crypto' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Crypto encode data

Description: Encode data to base64<br>Input params: (data []byte)<br>

```go
func CryptoEncodeData() {
	data := []byte("This is a sample data")

	encodedData, err := gouse.EncodeData(data)
	if err != nil {
		gouse.Println("Error encoding data:", err)
		return
	}
	gouse.Println("Raw data:", string(data))
	gouse.Println("Encoded data:", string(encodedData))
}
```

## 2. Crypto decode data

Description: Decode data from base64<br>Input params: (data []byte)<br>

```go
func CryptoDecodeData() {
	data := []byte("VGhpcyBpcyBhIHNhbXBsZSBkYXRh")
	decodedData, err := gouse.DecodeData(data)
	if err != nil {
		gouse.Println("Error decoding data:", err)
		return
	}
	gouse.Println("Raw encoded data:", string(data))
	gouse.Println("Decoded data:", string(decodedData))
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
		gouse.Println("Error encrypting data:", err)
		return
	}
	gouse.Println("Raw data:", string(data))
	gouse.Println("Encrypted data:", string(encryptedData))
}
```

## 6. Crypto compare password

Description: Compare string with the original password<br>Input params: (data string, password string)<br>

```go
func CryptoComparePassword() {
	data := "$2a$10$bcA002IOHi5SYHNH4lmIbuHjHplGl7TQZ.MznNrL1N70vAi7ovTa2"
	err := gouse.ComparePassword(data, "This is a sample data")
	if err != nil {
		gouse.Println("Error decrypting data:", err)
		return
	}
	println("Password matched")
}
```
