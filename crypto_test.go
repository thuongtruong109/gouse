package gouse

import (
	"bytes"
	"os"
	"testing"
)

func TestEncodeAndDecodeData(t *testing.T) {
	originalData := []byte("This is test data!")
	invalidData := []byte("Invalid base64 data!!!")

	encodedData, err := EncodeData(originalData)
	if err != nil {
		t.Errorf("EncodeData returned an error: %v", err)
	}

	if bytes.Equal(encodedData, originalData) {
		t.Errorf("Encoded data should not match the original data")
	}

	decodedData, err := DecodeData(encodedData)
	if err != nil {
		t.Errorf("DecodeData returned an error: %v", err)
	}

	if !bytes.Equal(decodedData, originalData) {
		t.Errorf("Decoded data does not match the original data")
	}

	_, err = DecodeData(invalidData)
	if err == nil {
		t.Errorf("DecodeData should fail with invalid base64 data")
	}
}

func TestEncryptAndDecryptFile(t *testing.T) {
	// Test data
	originalContent := []byte("This is the content of the file.")
	testFile := "mockdata/crypto_test.txt"
	password := []byte("strongpassword123")

	err := os.WriteFile(testFile, originalContent, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile)

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("EncryptFile panicked: %v", r)
			}
		}()
		EncryptFile(testFile, password)
	}()

	encryptedContent, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read encrypted file: %v", err)
	}
	if bytes.Equal(encryptedContent, originalContent) {
		t.Error("File content should be encrypted, but it matches the original content")
	}

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("DecryptFile panicked: %v", r)
			}
		}()
		DecryptFile(testFile, password)
	}()

	decryptedContent, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read decrypted file: %v", err)
	}
	if !bytes.Equal(decryptedContent, originalContent) {
		t.Error("Decrypted content does not match the original content")
	}
}

func TestEncryptAndComparePassword(t *testing.T) {
	originalPassword := "securePassword123"
	wrongPassword := "wrongPassword456"

	hashedPassword, err := EncryptPassword(originalPassword)
	if err != nil {
		t.Errorf("EncryptPassword returned an error: %v", err)
	}

	if hashedPassword == originalPassword {
		t.Errorf("Hashed password should not match the original password")
	}

	if err := ComparePassword(hashedPassword, originalPassword); err != nil {
		t.Errorf("ComparePassword failed with the correct password: %v", err)
	}

	if err := ComparePassword(hashedPassword, wrongPassword); err == nil {
		t.Errorf("ComparePassword should fail with the wrong password")
	}
}
