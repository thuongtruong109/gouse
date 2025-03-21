package samples

import "github.com/thuongtruong109/gouse"

func NumberClamp() {
	println("Clamp number: ", gouse.Clamp(5, 1, 10))
}

func NumberCheckInRange() {
	println("Check number is in range: ", gouse.InRange(5, 1, 10))
}
