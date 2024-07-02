package function

import "github.com/thuongtruong109/gouse/function"

func SampleFuncInterval() {
	function.Interval(func() {
		println("Interval")
	}, 1)
}
