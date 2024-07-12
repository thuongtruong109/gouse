package net

import (
	"fmt"

	"github.com/thuongtruong109/gouse/net"
)

/*
Title: Port Checker
Description: Check if a port is open
Package: api
Input: protocol, hostname, port
*/

func SampleApiPortChecker() {
	open := net.PortChecker("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
}
