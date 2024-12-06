package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Return a random number between min and max
Input params: (min, max int)
*/
func RandomNumber() {
	fmt.Println(gouse.RandNum(1, 100))
}

/*
Description: Return a random id (string) with current timestamp
*/
func RandomID() {
	fmt.Println(gouse.RandID())
}

/*
Description: Return a random string with n characters length
Input params: (length int)
*/
func RandomString() {
	fmt.Println(gouse.RandStr(10))
}

/*
Description: Return a random digit number with n characters length
Input params: (length int)
*/
func RandomDigit() {
	fmt.Println(gouse.RandDigit(10))
}

/*
Description: Return a random UUID (format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
*/
func RandomUUID() {
	fmt.Println(gouse.UUID())
}
