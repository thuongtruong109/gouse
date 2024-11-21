package gouse

func Clamp(value int, min int, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}

	return value
}

func InRange(value int, min int, max int) bool {
	return value >= min && value <= max
}
