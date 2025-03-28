
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Random' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Random number

Description: Return a random number between min and max<br>Input params: (min, max int)<br>

```go
func RandomNumber() {
	fmt.Println("Random number [1, 10): ", gouse.RandNum(1, 100))
}
```

## 2. Random i d

Description: Return a random id (string) with current timestamp<br>

```go
func RandomID() {
	fmt.Println("Generate random ID: ", gouse.RandID())
}
```

## 3. Random string

Description: Return a random string with n characters length<br>Input params: (length int)<br>

```go
func RandomString() {
	fmt.Println("Random string: ", gouse.RandStr(10))
}
```

## 4. Random digit

Description: Return a random digit number with n characters length<br>Input params: (length int)<br>

```go
func RandomDigit() {
	fmt.Println("Random digit: ", gouse.RandDigit(10))
}
```

## 5. Random u u i d

Description: Return a new random UUID (format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)<br>

```go
func RandomUUID() {
	fmt.Println("New uuid: ", gouse.UUID())
}
```
