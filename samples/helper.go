package samples

import "github.com/thuongtruong109/gouse"

func HelperRandomID() {
	println("Generate random ID: ", gouse.RandID())
}

func HelperUUID() {
	println("New uuid: ", gouse.UUID())
}

func HelperAutoMdDoc() {
	inputFilePath := "main.go"
	outputFilePath := "main.md"
	gouse.Go2Md(inputFilePath, outputFilePath)
}
