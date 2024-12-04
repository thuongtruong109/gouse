package samples

import "github.com/thuongtruong109/gouse"

func NumClamp() {
	println("Clamp number: ", gouse.Clamp(5, 1, 10))
}

func NumRandom() {
	random := gouse.RandNum(1, 10)
	println("Random number [1, 10): ", random)
}

func NumInRange() {
	println("Check number is in range: ", gouse.InRange(5, 1, 10))
}
