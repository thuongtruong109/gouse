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

func SampleNetCheckWithStatusCode() {
	statusCode, err := gouse.CheckHrefStatusCode("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Status code: ", statusCode)
}

func SampleNetConnectTime() {
	connectTime, err := gouse.HrefConnectTime("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connect time: %fs\n", connectTime)
}

func SampleNetEncode() {
	println("Encode: ", gouse.EncodeHref("https://google.com"))
}

func SampleNetDecode() {
	println("Decode: ", gouse.DecodeHref("https%3A%2F%2Fgoogle.com"))
}

func SampleNetHeader() {
	header, err := gouse.HrefHeader("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(gouse.MapAsString(header))
}

/*
Title: Port Checker
Description: Check if a port is open
Package: api
Input: protocol, hostname, port
*/

func SampleApiPortChecker() {
	open := gouse.CheckPort("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
}

/*
Title: Port Scanner
Description: This sample will scan for open ports on a given host.
Package: api
Input: protocol, hostname, start, end
*/

func SampleApiPortScanner() {
	gouse.ScanPort("tcp", "127.0.0.1", 3000, 8080)
}

func SampleNetProxy() {
	gouse.Proxy("5000", []string{"http://localhost:3000", "http://localhost:3001"})
}

func SampleNetOpen() {
	gouse.OpenHref("https://google.com")
}
