package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Encode data to base64
Input params: (data []byte)
*/
func CryptoEncodeData() {
	data := []byte("This is a sample data")

	encodedData, err := gouse.EncodeData(data)
	if err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}
	fmt.Println("Raw data:", string(data))
	fmt.Println("Encoded data:", string(encodedData))
}

/*
Description: Decode data from base64
Input params: (data []byte)
*/
func CryptoDecodeData() {
	data := []byte("VGhpcyBpcyBhIHNhbXBsZSBkYXRh")
	decodedData, err := gouse.DecodeData(data)
	if err != nil {
		fmt.Println("Error decoding data:", err)
		return
	}
	fmt.Println("Raw encoded data:", string(data))
	fmt.Println("Decoded data:", string(decodedData))
}

/*
Description: Encrypt data in file
Input params: (filename string, password []byte)
*/
func CryptoEncryptFile() {
	gouse.EncryptFile("sample.txt", []byte("password"))
	println("File content encrypted")
}

/*
Description: Decrypt data in file
Input params: (filename string, password []byte)
*/
func CryptoDecryptFile() {
	gouse.DecryptFile("sample.txt", []byte("password"))
	println("File content decrypted")
}

/*
Description: Encrypt password string
Input params: (data string)
*/
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

/*
Description: Compare string with the original password
Input params: (data string, password string)
*/
func CryptoComparePassword() {
	data := "$2a$10$bcA002IOHi5SYHNH4lmIbuHjHplGl7TQZ.MznNrL1N70vAi7ovTa2"
	err := gouse.ComparePassword(data, "This is a sample data")
	if err != nil {
		fmt.Println("Error decrypting data:", err)
		return
	}
	println("Password matched")
}
