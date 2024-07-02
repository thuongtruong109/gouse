package io

import (
	"fmt"

	"github.com/thuongtruong109/gouse/io"
)

func SampleIoZip() {
	filesToZip := []string{"file1.txt", "file2.txt"}
	zipFileName := "archive.zip"
	err := io.Zip(zipFileName, filesToZip)
	if err != nil {
		fmt.Println("Error zipping files:", err)
	}

	println("Files zipped successfully:", zipFileName)
}
