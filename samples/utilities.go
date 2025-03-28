package samples

import "github.com/thuongtruong109/gouse"

/*
Description: Auto generate markdown document from go source code
Input params: (inputFilePath: string, outputFilePath: string)
*/
func UtilsGoToMarkdown() {
	inputFilePath := "main.go"
	outputFilePath := "main.md"
	gouse.Go2Md(inputFilePath, outputFilePath)
}
