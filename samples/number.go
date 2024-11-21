package samples

import "github.com/thuongtruong109/gouse"

func SampleNumClamp() {
	println("Clamp number: ", gouse.Clamp(5, 1, 10))
}

func SampleNumRandom() {
	random := gouse.RandNum(1, 10)
	println("Random number [1, 10): ", random)
}

func SampleNumInRange() {
	println("Check number is in range: ", gouse.InRange(5, 1, 10))
}
