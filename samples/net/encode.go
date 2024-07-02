package net

import "github.com/thuongtruong109/gouse/net"

func SampleNetEncode() {
	println("Encode: ", net.Encode("https://google.com"))
}

func SampleNetDecode() {
	println("Decode: ", net.Decode("https%3A%2F%2Fgoogle.com"))
}
