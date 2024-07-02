package crypto

import "github.com/thuongtruong109/gouse/crypto"

func SampleCryptoEncryptFile() {
	crypto.EncryptFile("sample.txt", []byte("password"))
	println("File content encrypted")
}

func SampleCryptoDecryptFile() {
	crypto.DecryptFile("sample.txt", []byte("password"))
	println("File content decrypted")
}
