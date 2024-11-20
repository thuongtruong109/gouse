package gouse

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
)

/* Data */

func EncodeData(input []byte) ([]byte, error) {
	encodedData := base64.StdEncoding.EncodeToString(input)
	return []byte(encodedData), nil
}

func DecodeData(input []byte) ([]byte, error) {
	decodedData, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		return nil, err
	}
	return decodedData, nil
}

/* File */

func EncryptFile(source string, password []byte) {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}

	plaintext, err := os.ReadFile(source)
	if err != nil {
		panic(err.Error())
	}

	key := password
	nonce := make([]byte, 12)

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	ciphertext = append(ciphertext, nonce...)

	f, err := os.Create(source)
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		panic(err.Error())
	}
}

func DecryptFile(source string, password []byte) {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}

	ciphertext, err := os.ReadFile(source)
	if err != nil {
		panic(err.Error())
	}

	key := password
	salt := ciphertext[len(ciphertext)-12:]
	str := hex.EncodeToString(salt)

	nonce, err := hex.DecodeString(str)
	if err != nil {
		panic(err.Error())
	}

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext[:len(ciphertext)-12], nil)
	if err != nil {
		panic(err.Error())
	}

	f, err := os.Create(source)
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(plaintext))
	if err != nil {
		panic(err.Error())
	}
}

/* Password */

func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", ERROR_HASH_PASSWORD
	}

	return string(hashedPassword), nil
}

func DecryptPassword(password, receivedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(receivedPassword)); err != nil {
		return ERROR_COMPARE_PASSWORD
	}

	return nil
}
