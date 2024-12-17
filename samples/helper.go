package samples

import "github.com/thuongtruong109/gouse"

/*
Description: Generate a new UUID
*/
func UUID() {
	println("New uuid: ", gouse.UUID())
}

/*
Description: Auto generate markdown document from go source code
Input params: (inputFilePath: string, outputFilePath: string)
*/
func AutoMarkdownDocument() {
	inputFilePath := "main.go"
	outputFilePath := "main.md"
	gouse.Go2Md(inputFilePath, outputFilePath)
}
