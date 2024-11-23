package samples

import "github.com/thuongtruong109/gouse"

func SampleHelperRandomID() {
	println("Generate random ID: ", gouse.RandID())
}

func SampleHelperUUID() {
	println("New uuid: ", gouse.UUID())
}

func SampleHelperAutoMdDoc() {
	inputFilePath := "main.go"
	outputFilePath := "main.md"
	gouse.Go2Md(inputFilePath, outputFilePath)
}
