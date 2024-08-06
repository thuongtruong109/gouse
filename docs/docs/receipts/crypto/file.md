# File

## Imports

```go
import (
	"github.com/thuongtruong109/gouse/crypto")
```
## Functions


### SampleCryptoEncryptFile

```go
func SampleCryptoEncryptFile() {
	crypto.EncryptFile("sample.txt", []byte("password"))
	println("File content encrypted")
}```
## Imports

```go
import (
	"github.com/thuongtruong109/gouse/crypto")
```
## Functions


### SampleCryptoDecryptFile

```go
func SampleCryptoDecryptFile() {
	crypto.DecryptFile("sample.txt", []byte("password"))
	println("File content decrypted")
}```
