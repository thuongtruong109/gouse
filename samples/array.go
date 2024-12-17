package samples

import "github.com/thuongtruong109/gouse"

/*
Description: Chunk an array into smaller arrays of a specified size
Input params: (array, size)
*/
func ArrayChunk() {
	gouse.Println("[int]: ", gouse.Chunk([]int{1, -2, 3, -4, 5, 6}, 3))
	gouse.Println("[uint]: ", gouse.Chunk([]uint{1, 2, 3, 4, 5, 6}, 3))
	gouse.Println("[float]: ", gouse.Chunk([]float64{1.2, 2.3, 3.4, 4.5, 5.6, 6.7}, 3))
	gouse.Println("[string]: ", gouse.Chunk([]string{"1", "2", "3", "4", "5", "6"}, 3))
	gouse.Println("[rune]: ", gouse.Chunk([]rune{'a', 'b', 'c', 'd', 'e', 'f'}, 3))
	gouse.Println("[complex]: ", gouse.Chunk([]complex128{1 + 2i, 2 + 3i, 3 + 4i, 4 + 5i, 5 + 6i, 6 + 7i}, 3))
	gouse.Println("[struct]: ", gouse.Chunk([]struct{ a int }{{1}, {2}, {3}, {4}, {5}, {6}}, 3))
}

/*
Description: Removing falsy values (false, null, 0, "", undefined, and NaN) from an array
Input params: (array)
*/
func ArrayCompact() {
	result := gouse.Compact([]interface{}{1, -2, 3, -4, 5, 6, 0, 0.0, "", false, nil})
	gouse.Println("Compact remove all falsy values: ", result)
}

/*
Description: Check difference items between two arrays
Input params: (array1, array2)
*/
func ArrayDifference() {
	gouse.Println("[int]: ", gouse.Diff([]int{1, -2, 3, -4, 5, 6}, []int{1, 2, 3, 4, 5, 6}))
	gouse.Println("[uint]: ", gouse.Diff([]uint{1, 2, 3, 4, 5, 7}, []uint{1, 2, 3, 4, 5, 6}))
	gouse.Println("[float]: ", gouse.Diff([]float64{1.2, 2.3, 3.4}, []float64{4.5, 5.6, 6.7}))
	gouse.Println("[string]: ", gouse.Diff([]string{"1", "4", "5", "6"}, []string{"1", "2", "3", "6"}))
	gouse.Println("[rune]: ", gouse.Diff([]rune{'a', 'b', 'd', 'e', 'f'}, []rune{'a', 'b', 'c', 'f'}))
	gouse.Println("[complex]: ", gouse.Diff([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, []complex128{1 + 2i, 2 + 3i, 6 + 7i}))
	gouse.Println("[struct]: ", gouse.Diff([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, []struct{ a int }{{1}, {2}}))
}

/*
Description: Drop n elements in array (default n = 1)
Input params: (array, n)
*/
func ArrayDrop() {
	gouse.Println("[int] with default: ", gouse.Drop([]int{1, -2, 3, -4, 5, 6}))
	gouse.Println("[int]: ", gouse.Drop([]int{1, -2, 3, -4, 5, 6}, 2))
	gouse.Println("[uint]: ", gouse.Drop([]uint{1, 2, 3, 4, 5, 7}, 2))
	gouse.Println("[float]: ", gouse.Drop([]float64{1.2, 2.3, 3.4}, 2))
	gouse.Println("[string]: ", gouse.Drop([]string{"1", "4", "5", "6"}, 2))
	gouse.Println("[rune]: ", gouse.Drop([]rune{'a', 'b', 'd', 'e', 'f'}, 2))
	gouse.Println("[complex]: ", gouse.Drop([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, 2))
	gouse.Println("[struct]: ", gouse.Drop([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, 2))
}

/*
Description: Compare is equal between two elements
Input params: (element1, element2)
*/
func ArrayEqual() {
	println("[int]: ", gouse.Equal(1, 1))
	println("[uint]: ", gouse.Equal(uint(1), uint(1)))
	println("[float]: ", gouse.Equal(1.2, 1.1))
	println("[string]: ", gouse.Equal("1", "0"))
	println("[rune]: ", gouse.Equal('a', 'a'))
	println("[bool]: ", gouse.Equal(true, true))
	println("[complex]: ", gouse.Equal(1+2i, 1+2i))
	println("[struct]: ", gouse.Equal(struct{ a int }{1}, struct{ a int }{1}))
}

/*
Description: Filter elements in array by pass condition in callback function
Input params: (array, callback)
*/
func ArrayFilterBy() {
	println("[int]: ", gouse.FilterBy([]int{1, -2, 3, -4, 5, 6}, func(v int) bool {
		return v > 2
	}))

	println("[uint]: ", gouse.FilterBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) bool {
		return v > 2
	}))

	println("[float]: ", gouse.FilterBy([]float64{1.2, 2.3, 3.4}, func(v float64) bool {
		return v > 2
	}))

	println("[string]: ", gouse.FilterBy([]string{"1", "4", "5", "6"}, func(v string) bool {
		return v > "3"
	}))

	println("[rune]: ", gouse.FilterBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) bool {
		return v > 'c'
	}))

	println("[complex]: ", gouse.FilterBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) bool {
		return real(v) > 3
	}))

	println("[struct]: ", gouse.FilterBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) bool {
		return v.a > 0
	}))
}

/*
Description: Find element in array by pass condition in callback function
Input params: (array, callback)
*/
func ArrayFindBy() {
	println("[int]: ", gouse.FindBy([]int{1, -2, 3, -4, 5, 6}, func(v int) bool {
		return v == 3
	}))

	println("[uint]: ", gouse.FindBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) bool {
		return v == 3
	}))

	println("[float]: ", gouse.FindBy([]float64{1.2, 2.3, 3.4}, func(v float64) bool {
		return v == 3.4
	}))

	println("[string]: ", gouse.FindBy([]string{"1", "4", "5", "6"}, func(v string) bool {
		return v == "5"
	}))

	println("[rune]: ", gouse.FindBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) bool {
		return v == 'e'
	}))

	println("[complex]: ", gouse.FindBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) bool {
		return v == 5+6i
	}))

	gouse.Println("[struct]: ", gouse.FindBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) bool {
		return v.a == 3
	}))
}

/*
Description: Loop array then handler with callback function
Input params: (array, callback)
*/
func ArrayForBy() {
	print("[int]: ")
	gouse.ForBy([]int{1, -2, 3, -4, 5, 6}, func(v int) {
		println(v)
	})

	print("[uint]: ")
	gouse.ForBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) {
		println(v)
	})

	print("[float]: ")
	gouse.ForBy([]float64{1.2, 2.3, 3.4}, func(v float64) {
		println(v)
	})

	print("[string]: ")
	gouse.ForBy([]string{"1", "4", "5", "6"}, func(v string) {
		println(v)
	})

	print("[rune]: ")
	gouse.ForBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) {
		println(v)
	})

	print("[complex]: ")
	gouse.ForBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) {
		println(v)
	})

	print("[struct]: ")
	gouse.ForBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) {
		gouse.Println(v)
	})
}

/*
Description: Check element is exist in array
Input params: (array, element)
*/
func ArrayIncludes() {
	println("[int]: ", gouse.IncludesArr([]int{1, -2, 3}, 1))
	println("[uint]: ", gouse.IncludesArr([]uint{1, 2, 3}, 1))
	println("[float]: ", gouse.IncludesArr([]float64{1.2, 2.3, 3.4}, 1.2))
	println("[string]: ", gouse.IncludesArr([]string{"1", "2", "3"}, "0"))
	println("[rune]: ", gouse.IncludesArr([]rune{'a', 'b', 'c'}, 'a'))
	println("[complex]: ", gouse.IncludesArr([]complex128{1 + 2i, 2 + 3i}, 1+2i))
	println("[struct]: ", gouse.IncludesArr([]struct{ a int }{{1}, {2}}, struct{ a int }{3}))
}

/*
Description: Get index of element in array
Input params: (array, element)
*/
func ArrayIndex() {
	println("[int]: ", gouse.IndexOfArr([]int{1, -2, 3, -4, 5, 6}, 3))
	println("[uint]: ", gouse.IndexOfArr([]uint{1, 2, 3, 4, 5, 7}, 3))
	println("[float]: ", gouse.IndexOfArr([]float64{1.2, 2.3, 3.4}, 3.4))
	println("[string]: ", gouse.IndexOfArr([]string{"1", "4", "5", "6"}, "5"))
	println("[rune]: ", gouse.IndexOfArr([]rune{'a', 'b', 'd', 'e', 'f'}, 'e'))
	println("[complex]: ", gouse.IndexOfArr([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, 5+6i))
	println("[struct]: ", gouse.IndexOfArr([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, struct{ a int }{3}))
}

/*
Description: Find index of element pass condition in callback function
Input params: (array, callback)
*/
func ArrayIndexBy() {
	println("[int]: ", gouse.IndexBy([]int{1, -2, 3, -4, 5, 6}, func(v int) bool {
		return v == 3
	}))

	println("[uint]: ", gouse.IndexBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) bool {
		return v == 3
	}))

	println("[float]: ", gouse.IndexBy([]float64{1.2, 2.3, 3.4}, func(v float64) bool {
		return v == 3.4
	}))

	println("[string]: ", gouse.IndexBy([]string{"1", "4", "5", "6"}, func(v string) bool {
		return v == "5"
	}))

	println("[rune]: ", gouse.IndexBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) bool {
		return v == 'e'
	}))

	println("[complex]: ", gouse.IndexBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) bool {
		return v == 5+6i
	}))

	println("[struct]: ", gouse.IndexBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) bool {
		return v.a == 3
	}))
}

/*
Description: Intersection arrays
Input params: (array1, array2)
*/
func ArrayIntersect() {
	println("[int]: ", gouse.Intersect([]int{1, -2, 3, -4, 5, 6}, []int{1, 2, 3, 4, 5, 6}))
	println("[uint]: ", gouse.Intersect([]uint{1, 2, 3, 4, 5, 7}, []uint{1, 2, 3, 4, 5, 6}))
	println("[float]: ", gouse.Intersect([]float64{1.2, 2.3, 3.4}, []float64{1.2, 4.5, 5.6, 6.7}))
	println("[string]: ", gouse.Intersect([]string{"1", "4", "5", "6"}, []string{"1", "2", "3", "6"}))
	println("[rune]: ", gouse.Intersect([]rune{'a', 'b', 'd', 'e', 'f'}, []rune{'a', 'b', 'c', 'f'}))
	println("[complex]: ", gouse.Intersect([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, []complex128{1 + 2i, 2 + 3i, 6 + 7i}))
	println("[struct]: ", gouse.Intersect([]struct{ a int }{{1}, {-2}, {3}, {4}, {5}, {6}}, []struct{ a int }{{1}, {2}}))
}

/*
Description: Find key of element pass condition in callback function
Input params: (array, callback)
*/
func ArrayKeyBy() {
	println("[int]: ", gouse.KeyBy([]int{1, -2, 3, -4, 5, 6}, func(v int) bool {
		return v == 3
	}))

	println("[uint]: ", gouse.KeyBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) bool {
		return v == 3
	}))

	println("[float]: ", gouse.KeyBy([]float64{1.2, 2.3, 3.4}, func(v float64) bool {
		return v == 3.4
	}))

	println("[string]: ", gouse.KeyBy([]string{"1", "4", "5", "6"}, func(v string) bool {
		return v == "5"
	}))

	println("[rune]: ", gouse.KeyBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) bool {
		return v == 'e'
	}))

	println("[complex]: ", gouse.KeyBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) bool {
		return v == 5+6i
	}))

	println("[struct]: ", gouse.KeyBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) bool {
		return v.a == 3
	}))
}

/*
Description: Map array then handler with callback function
Input params: (array, callback)
*/
func ArrayMapBy() {
	gouse.Println("[int]: ", gouse.MapBy([]int{1, -2, 3, -4, 5, 6}, func(v int) int {
		return v * 2
	}))

	gouse.Println("[uint]: ", gouse.MapBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) uint {
		return v * 2
	}))

	gouse.Println("[float]: ", gouse.MapBy([]float64{1.2, 2.3, 3.4}, func(v float64) float64 {
		return v * 2
	}))

	gouse.Println("[string]: ", gouse.MapBy([]string{"1", "4", "5", "6"}, func(v string) string {
		return v + "1"
	}))

	gouse.Println("[rune]: ", gouse.MapBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) rune {
		return v + 1
	}))

	gouse.Println("[complex]: ", gouse.MapBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) complex128 {
		return v * 2
	}))

	gouse.Println("[struct]: ", gouse.MapBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) struct{ a int } {
		return struct{ a int }{v.a * 2}
	}))
}

/*
Description: Find min element in array
Input params: (array)
*/
func ArrayMin() {
	println("[int]: ", gouse.MinArr([]int{1, -2, 3}))
	println("[uint]: ", gouse.MinArr([]uint{1, 2, 3}))
	println("[string]: ", gouse.MinArr([]string{"z", "d", "m"}))
	println("[rune]: ", string(gouse.MinArr([]rune{'z', 'd', 'm'})))
	println("[float]: ", gouse.MinArr([]float64{1.2, 2.3, 3.4}))
}

/*
Description: Find max element in array
Input params: (array)
*/
func ArrayMax() {
	println("[int]: ", gouse.MaxArr([]int{1, -2, 3}))
	println("[uint]: ", gouse.MaxArr([]uint{1, 2, 3}))
	println("[string]: ", gouse.MaxArr([]string{"z", "d", "m"}))
	println("[rune]: ", string(gouse.MaxArr([]rune{'z', 'd', 'm'})))
	println("[float]: ", gouse.MaxArr([]float64{1.2, 2.3, 3.4}))
}

/*
Description: Find most frequency element in array
Input params: (array)
*/
func ArrayMost() {
	println("[int]: ", gouse.Most([]int{1, -2, 3, 2, 2, 1, 2, 3}))
	println("[uint]: ", gouse.Most([]uint{1, 2, 3, 2, 2, 1, 2, 3}))
	gouse.Println("[float]: ", gouse.Most([]float64{1.2, 2.3, 3.4, 2.3, 2.3, 1.2, 2.3, 3.4}))
	println("[string]: ", gouse.Most([]string{"1", "2", "3", "2", "2", "1", "2", "3"}))
	println("[rune]: ", string(gouse.Most([]rune{'a', 'b', 'c', 'b', 'b', 'a', 'b', 'c'})))
	gouse.Println("[complex]: ", gouse.Most([]complex128{1 + 2i, 2 + 3i, 3 + 4i, 2 + 3i, 2 + 3i, 1 + 2i, 2 + 3i, 3 + 4i}))
	gouse.Println("[struct]: ", gouse.Most([]struct{ a int }{{1}, {2}, {3}, {2}, {2}, {1}, {2}, {3}}))
}

/*
Description: Merge arrays
Input params: (array1, array2, array3, ...)
*/
func ArrayMerge() {
	println("[int]: ", gouse.Merge([]int{1, -2, 3, -4, 5, 6}, []int{1, 2, 3, 4, 5, 6}, []int{1, -2, 3, -4, 5, 6}))
	println("[uint]: ", gouse.Merge([]uint{1, 2, 3, 4, 5, 7}, []uint{1, 2, 3, 4, 5, 6}))
	println("[float]: ", gouse.Merge([]float64{1.2, 2.3, 3.4}, []float64{1.2, 4.5, 5.6, 6.7}))
	println("[string]: ", gouse.Merge([]string{"1", "4", "5", "6"}, []string{"1", "2", "3", "6"}))
	println("[rune]: ", gouse.Merge([]rune{'a', 'b', 'd', 'e', 'f'}, []rune{'a', 'b', 'c', 'f'}))
	println("[complex]: ", gouse.Merge([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, []complex128{1 + 2i, 2 + 3i, 6 + 7i}))
	println("[struct]: ", gouse.Merge([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, []struct{ a int }{{1}, {2}}))
}

/*
Description: Remove element in array by pass condition in callback function
Input params: (array, callback)
*/
func ArrayRejectBy() {
	println("[int]: ", gouse.RejectBy([]int{1, -2, 3, -4, 5, 6}, func(v int) bool {
		return v > 2
	}))

	println("[uint]: ", gouse.RejectBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) bool {
		return v > 2
	}))

	println("[float]: ", gouse.RejectBy([]float64{1.2, 2.3, 3.4}, func(v float64) bool {
		return v > 2
	}))

	println("[string]: ", gouse.RejectBy([]string{"1", "4", "5", "6"}, func(v string) bool {
		return v > "3"
	}))

	println("[rune]: ", gouse.RejectBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) bool {
		return v > 'c'
	}))

	println("[complex]: ", gouse.RejectBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) bool {
		return real(v) > 3
	}))
}

/*
Description: Calculate sum of elements in array
Input params: (array)
*/
func ArraySum() {
	println("[int]: ", gouse.SumArr([]int{1, -2, 3}))
	println("[uint]: ", gouse.SumArr([]uint{1, 2, 3}))
	gouse.Println("[float]: ", gouse.SumArr([]float64{1.2, 2.3, 3.4}))
	println("[rune]: ", gouse.SumArr([]rune{'a', 'b', 'c'}))
	gouse.Println("[complex]: ", gouse.SumArr([]complex128{1 + 2i, 2 + 3i, 3 + 4i}))
}
