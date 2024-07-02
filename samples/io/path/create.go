package path

import (
	"fmt"

	"github.com/thuongtruong109/gouse/io"
)

func SampleIoCreatePath() {
	relativePath := "tmp/example.txt"

	if err := io.CreatePath(relativePath); err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	println("File created successfully.")
}
