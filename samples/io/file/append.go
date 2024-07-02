package file

import "github.com/thuongtruong109/gouse/io"

func SampleIoAppendToFile() {
	err := io.AppendFile("data.json", []string{"this is data 3", "this is data 4"})
	if err != nil {
		println(err.Error())
	}
	println("file appended")
}
