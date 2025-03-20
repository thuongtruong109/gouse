package samples

import "github.com/thuongtruong109/gouse"

/*
Description: Check if a URL status
Input params: (url string)
*/
func CheckURL() {
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
func ConnectTime() {
	connectTime, err := gouse.ConTimeUrl("https://google.com")
	if err != nil {
		panic(err)
	}

	gouse.Printf("Connect time: %fs\n", connectTime)
}

/*
Description: Encode a URL
Input params: (url string)
*/
func EncodeURL() {
	println("Encode: ", gouse.EncodeUrl("https://google.com"))
}

/*
Description: Decode a URL
Input params: (url string)
*/
func DecodeURL() {
	println("Decode: ", gouse.DecodeUrl("https%3A%2F%2Fgoogle.com"))
}

/*
Description: Get the header of a URL
Input params: (url string)
*/
func HeaderURL() {
	header, err := gouse.HeaderUrl("https://google.com")
	if err != nil {
		panic(err)
	}

	gouse.Println(gouse.MapAsString(header))
}

/*
Description: Check if a port is open
Input params: (protocol, hostname string, port int)
*/
func PortChecker() {
	open := gouse.CheckPort("tcp", "localhost", 1313)
	gouse.Printf("Port Open: %t\n", open)
}

/*
Description: Scan for open ports on a given host.
Input params: (protocol, hostname string, startPort, endPort int)
*/
func PortScanner() {
	gouse.ScanPort("tcp", "127.0.0.1", 3000, 8080)
}

/*
Description: Proxy wrapper to another port
Input params: (port string, urls []string{})
*/
func Proxy() {
	gouse.Proxy("5000", []string{"http://localhost:3000", "http://localhost:3001"})
}

/*
Description: Open a URL in the default browser
Input params: (url string)
*/
func OpenURL() {
	gouse.OpenUrl("https://google.com")
}
