package function

import "github.com/thuongtruong109/gouse/function"

func SampleFuncRetry() {
	function.Retry(func() error {
		println("Retry")
		return nil
	}, 3, 1)
}
