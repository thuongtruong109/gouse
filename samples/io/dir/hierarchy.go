package dir

import "github.com/thuongtruong109/gouse/io"

func SampleIoHierarchyDir() {
	data, err := io.HierarchyDir(".")
	if err != nil {
		println(err.Error())
		return
	}

	for _, v := range data {
		println(v)
	}
}
