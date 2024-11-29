package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

func SampleNetCheck() {
	ok, err := gouse.CheckHref("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Response: ", ok)
}

/*
Description: Check if a URL is valid
Input params: (url)
*/
func SampleNetCheckWithStatusCode() {
	statusCode, err := gouse.CheckHrefStatusCode("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Status code: ", statusCode)
}

/*
Description: Calculate the time it takes to connect to a URL
Input params: (url)
*/
func SampleNetConnectTime() {
	connectTime, err := gouse.HrefConnectTime("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connect time: %fs\n", connectTime)
}

/*
Description: Encode a URL
Input params: (url)
*/
func SampleNetEncode() {
	println("Encode: ", gouse.EncodeHref("https://google.com"))
}

/*
Description: Decode a URL
Input params: (url)
*/
func SampleNetDecode() {
	println("Decode: ", gouse.DecodeHref("https%3A%2F%2Fgoogle.com"))
}

/*
Description: Get the header of a URL
Input params: (url)
*/
func SampleNetHeader() {
	header, err := gouse.HrefHeader("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(gouse.MapAsString(header))
}

/*
Description: Check if a port is open
Input params: (protocol, hostname, port)
*/
func SampleApiPortChecker() {
	open := gouse.CheckPort("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
}

/*
Description: Scan for open ports on a given host.
Input params: (protocol, hostname, start port, end port)
*/
func SampleApiPortScanner() {
	gouse.ScanPort("tcp", "127.0.0.1", 3000, 8080)
}

/*
Description: Proxy wrapper to another port
Input params: (port, []string{urls})
*/
func SampleNetProxy() {
	gouse.Proxy("5000", []string{"http://localhost:3000", "http://localhost:3001"})
}

/*
Description: Open a URL in the default browser
Input params: (url)
*/
func SampleNetOpen() {
	gouse.OpenHref("https://google.com")
}
