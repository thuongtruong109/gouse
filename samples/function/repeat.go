package function

import "github.com/thuongtruong109/gouse/function"

func SampleFuncTimes() {
	function.Times(func() {
		println("Times")
	}, 3)
}
