package main

import (
	"flag"

	"github.com/thuongtruong109/gouse/samples"
	"github.com/thuongtruong109/gouse/strings"
)

func Starter() {
	println(strings.Capitalize("hello world from Gouse!"))
}

func main() {
	isDevFlag := flag.Bool("isDev", false, "toggle enviroment mode")
	flag.Parse()
	if *isDevFlag {
		samples.Run()
	} else {
		Starter()
	}
}
