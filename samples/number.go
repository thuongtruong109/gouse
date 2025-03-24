package samples

import "github.com/thuongtruong109/gouse"

/*
Description: Clamp a number between a minimum and maximum value
Input params: (number, min, max int) -> int
*/
func NumberClamp() {
	println("Clamp number: ", gouse.Clamp(5, 1, 10))
}

/*
Description: Check if a number is in a range
Input params: (number, min, max int) -> bool
*/
func NumberInRange() {
	println("Check number is in range: ", gouse.InRange(5, 1, 10))
}

/*
Description: Sort a list of numbers
Input params: (nums []int) -> []int
*/
func NumberSort() {
	nums := []int{5, 3, 8, 1, 2}
	println("Sort numbers: ", gouse.SortNum(nums))
}
