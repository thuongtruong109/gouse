package operator

import "github.com/thuongtruong109/gouse/math"

func SampleMathMin() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Min number: ", math.Min(num1, num2, num3, num4))
}