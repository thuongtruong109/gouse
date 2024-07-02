package file

import "github.com/thuongtruong109/gouse/io"

func SampleIoWriteToFile() {
	err := io.WriteFile("data.json", []string{"this is data 1", "this is data 2"})
	if err != nil {
		println(err.Error())
	}
	println("file written")
}
