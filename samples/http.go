package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Check if a URL status
Input params: (url string)
*/
func HttpCheckUrl() {
	isOk, statusCode, err := gouse.CheckUrl("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Response: ", isOk)
	println("Status code: ", statusCode)
}

/*
Description: Calculate the time it takes to connect to a URL
Input params: (url string)
*/
func HttpConnectTime() {
	connectTime, err := gouse.ConTimeUrl("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connect time: %fs\n", connectTime)
}

/*
Description: Encode a URL
Input params: (url string)
*/
func HttpEncodeUrl() {
	println("Encode: ", gouse.EncodeUrl("https://google.com"))
}

/*
Description: Decode a URL
Input params: (url string)
*/
func HttpDecodeUrl() {
	println("Decode: ", gouse.DecodeUrl("https%3A%2F%2Fgoogle.com"))
}

/*
Description: Get the header of a URL
Input params: (url string)
*/
func HttpHeaderUrl() {
	header, err := gouse.HeaderUrl("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(gouse.Map2Str(header))
}

/*
Description: Check if a port is open
Input params: (protocol, hostname string, port int)
*/
func HttpPortChecker() {
	open := gouse.CheckPort("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
}

/*
Description: Scan for open ports on a given host.
Input params: (protocol, hostname string, startPort, endPort int)
*/
func HttpPortScanner() {
	gouse.ScanPort("tcp", "127.0.0.1", 3000, 8080)
}

/*
Description: Open a URL in the default browser
Input params: (url string)
*/
func HttpOpenUrl() {
	gouse.OpenUrl("https://google.com")
}
