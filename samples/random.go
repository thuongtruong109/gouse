package samples

import "github.com/thuongtruong109/gouse"

/*
Description: Return a random number between min and max
Input params: (min, max int)
*/
func RandomNumber() {
	gouse.Println("Random number [1, 10): ", gouse.RandNum(1, 100))
}

/*
Description: Return a random id (string) with current timestamp
*/
func RandomID() {
	gouse.Println("Generate random ID: ", gouse.RandID())
}

/*
Description: Return a random string with n characters length
Input params: (length int)
*/
func RandomString() {
	gouse.Println(gouse.RandStr(10))
}

/*
Description: Return a random digit number with n characters length
Input params: (length int)
*/
func RandomDigit() {
	gouse.Println(gouse.RandDigit(10))
}

/*
Description: Return a random UUID (format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
*/
func RandomUUID() {
	gouse.Println(gouse.UUID())
}
