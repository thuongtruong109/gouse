package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

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
Description: Find intersection elements among arrays
Input params: (array1, array2, array3, ...)
*/
func ArrayIntersect() {
	fmt.Println(gouse.Intersect([]int{1, 2, 3}, []int{2, 3, 4}))
	fmt.Println(gouse.Intersect([]string{"a", "b", "c"}, []string{"b", "c", "d"}))
	fmt.Println(gouse.Intersect([]float64{1.1, 2.2, 3.3}, []float64{2.2, 3.3, 4.4}))
	println("[int]: ", gouse.Intersect([]int{1, -2, 3, -4, 5, 6}, []int{1, 2, 3, 4, 5, 6}))
	println("[uint]: ", gouse.Intersect([]uint{1, 2, 3, 4, 5, 7}, []uint{1, 2, 3, 4, 5, 6}))
	println("[float]: ", gouse.Intersect([]float64{1.2, 2.3, 3.4}, []float64{1.2, 4.5, 5.6, 6.7}))
	println("[string]: ", gouse.Intersect([]string{"1", "4", "5", "6"}, []string{"1", "2", "3", "6"}))
	println("[rune]: ", gouse.Intersect([]rune{'a', 'b', 'd', 'e', 'f'}, []rune{'a', 'b', 'c', 'f'}))
	println("[complex]: ", gouse.Intersect([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, []complex128{1 + 2i, 2 + 3i, 6 + 7i}))
	println("[struct]: ", gouse.Intersect([]struct{ a int }{{1}, {-2}, {3}, {4}, {5}, {6}}, []struct{ a int }{{1}, {2}}))
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
Description: Compare is equal between two elements
Input params: (element1, element2)
*/
func ArrayEqual() {
	fmt.Println("[int]:", gouse.Equal([]int{1}, []int{1}))
	fmt.Println("[int]:", gouse.Equal([]int{1}, []int{2}))

	fmt.Println("[uint]:", gouse.Equal([]uint{1}, []uint{1}))
	fmt.Println("[uint]:", gouse.Equal([]uint{1}, []uint{2}))

	fmt.Println("[float]:", gouse.Equal([]float64{1.2}, []float64{1.2}))
	fmt.Println("[float]:", gouse.Equal([]float64{1.2}, []float64{1.1}))

	fmt.Println("[string]:", gouse.Equal([]string{"1"}, []string{"1"}))
	fmt.Println("[string]:", gouse.Equal([]string{"1"}, []string{"0"}))

	fmt.Println("[rune]:", gouse.Equal([]rune{'a'}, []rune{'a'}))
	fmt.Println("[rune]:", gouse.Equal([]rune{'a'}, []rune{'b'}))

	fmt.Println("[bool]:", gouse.Equal([]bool{true}, []bool{true}))
	fmt.Println("[bool]:", gouse.Equal([]bool{true}, []bool{false}))

	fmt.Println("[complex]:", gouse.Equal([]complex128{1 + 2i}, []complex128{1 + 2i}))
	fmt.Println("[complex]:", gouse.Equal([]complex128{1 + 2i}, []complex128{2 + 3i}))

	type MyStruct struct{ a int }
	fmt.Println("[struct]:", gouse.Equal([]MyStruct{{1}}, []MyStruct{{1}}))
	fmt.Println("[struct]:", gouse.Equal([]MyStruct{{1}}, []MyStruct{{2}}))
}

/*
Description: Find unique elements in array
Input params: (array)
*/
func ArrayUnique() {
	println("[int]: ", gouse.UniqueArr([]int{1, -2, 3, 1, -2, 3}))
	println("[uint]: ", gouse.UniqueArr([]uint{1, 2, 3, 1, 2, 3}))
	println("[float]: ", gouse.UniqueArr([]float64{1.2, 2.3, 3.4, 1.2, 2.3, 3.4}))
	println("[string]: ", gouse.UniqueArr([]string{"1", "4", "5", "1", "4", "5"}))
	println("[rune]: ", gouse.UniqueArr([]rune{'a', 'b', 'd', 'a', 'b', 'd'}))
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
Description: Find least frequency element in array
Input params: (array)
*/
func ArrayLeast() {
	println("[int]: ", gouse.Least([]int{1, -2, 3, 2, 2, 1, 2, 3}))
	println("[uint]: ", gouse.Least([]uint{1, 2, 3, 2, 2, 1, 2, 3}))
	gouse.Println("[float]: ", gouse.Least([]float64{1.2, 2.3, 3.4, 2.3, 2.3, 1.2, 2.3, 3.4}))
	println("[string]: ", gouse.Least([]string{"1", "2", "3", "2", "2", "1", "2", "3"}))
	println("[rune]: ", string(gouse.Least([]rune{'a', 'b', 'c', 'b', 'b', 'a', 'b', 'c'})))
	gouse.Println("[complex]: ", gouse.Least([]complex128{1 + 2i, 2 + 3i, 3 + 4i, 2 + 3i, 2 + 3i, 1 + 2i, 2 + 3i, 3 + 4i}))
	gouse.Println("[struct]: ", gouse.Least([]struct{ a int }{{1}, {2}, {3}, {2}, {2}, {1}, {2}, {3}}))
}

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
Description: Fill array with value from start to end index
Input params: (array, value, start, end)
*/
func ArrayFill() {
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Original int array:", intArr)

	updatedIntArr := gouse.Fill(intArr, 0, 2, 6)
	fmt.Println("Updated int array:", updatedIntArr)

	strArr := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Println("\nOriginal string array:", strArr)

	updatedStrArr := gouse.Fill(strArr, "fruit", 1, 4)
	fmt.Println("Updated string array:", updatedStrArr)
}

/*
Description: Shift array to left
Input params: (array)
*/
func ArrayShift() {
	intArr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original int array:", intArr)

	shiftedIntArr := gouse.Shift(intArr)
	fmt.Println("Shifted int array:", shiftedIntArr)

	strArr := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Println("\nOriginal string array:", strArr)

	shiftedStrArr := gouse.Shift(strArr)
	fmt.Println("Shifted string array:", shiftedStrArr)
}

/*
Description: Unshift array to right
Input params: (array, values)
*/
func ArrayUnshift() {
	intArr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original int array:", intArr)

	shiftedIntArr := gouse.Unshift(intArr, 0)
	fmt.Println("Shifted int array:", shiftedIntArr)

	strArr := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Println("\nOriginal string array:", strArr)

	shiftedStrArr := gouse.Unshift(strArr, "fruit")
	fmt.Println("Shifted string array:", shiftedStrArr)
}

/*
Description: Pop n elements in array (default n = 1)
Input params: (array, n)
*/
func ArrayPop() {
	intArr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original int array:", intArr)

	poppedIntArr := gouse.Pop(intArr)
	fmt.Println("Popped int array:", poppedIntArr)

	strArr := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Println("\nOriginal string array:", strArr)

	poppedStrArr := gouse.Pop(strArr)
	fmt.Println("Popped string array:", poppedStrArr)
}

/*
Description: Push element to array
Input params: (array, values)
*/
func ArrayPush() {
	intArr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original int array:", intArr)

	pushedIntArr := gouse.Push(intArr, 6)
	fmt.Println("Pushed int array:", pushedIntArr)

	strArr := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Println("\nOriginal string array:", strArr)

	pushedStrArr := gouse.Push(strArr, "fig")
	fmt.Println("Pushed string array:", pushedStrArr)
}

/*
Description: Get elements from start to end index
Input params: (array, start, end)
*/
func ArraySlice() {
	intArr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original int array:", intArr)

	slicedIntArr, err := gouse.SliceArr(intArr, 2, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Sliced int array:", slicedIntArr)

	strArr := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Println("\nOriginal string array:", strArr)

	slicedStrArr, err := gouse.SliceArr(strArr, 1, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Sliced string array:", slicedStrArr)
}

/*
Description: Remove items from start index with n elements
Input params: (array, start, n)
*/
func ArraySplice() {
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Original int array:", intArr)

	splicedIntArr := gouse.SpliceArr(intArr, 3, 2, 99, 100)
	fmt.Println("Spliced int array:", splicedIntArr)

	strArr := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Println("\nOriginal string array:", strArr)

	splicedStrArr := gouse.SpliceArr(strArr, 1, 2, "fruit", "grape")
	fmt.Println("Spliced string array:", splicedStrArr)

	removeOnlyArr := gouse.SpliceArr(strArr, 0, 2)
	fmt.Println("\nArray after removing elements:", removeOnlyArr)
}

/*
Description: Take n elements from array
Input params: (array, n)
*/
func ArrayTake() {
	intArr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original int array:", intArr)

	takenIntArr := gouse.Take(intArr, 3)
	fmt.Println("Taken int array:", takenIntArr)

	strArr := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Println("\nOriginal string array:", strArr)

	takenStrArr := gouse.Take(strArr, 2)
	fmt.Println("Taken string array:", takenStrArr)
}

/*
Description: Reverse array
Input params: (array)
*/
func ArrayReverse() {
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Original int array:", intArr)

	reversedIntArr := gouse.ReverseArr(intArr)
	fmt.Println("Reversed int array:", reversedIntArr)

	strArr := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Println("\nOriginal string array:", strArr)

	reversedStrArr := gouse.ReverseArr(strArr)
	fmt.Println("Reversed string array:", reversedStrArr)

	floatArr := []float64{3.14, 2.71, 1.41, 0.577, 1.618}
	fmt.Println("\nOriginal float array:", floatArr)

	reversedFloatArr := gouse.ReverseArr(floatArr)
	fmt.Println("Reversed float array:", reversedFloatArr)
}

/*
Description: Find index of element in array
Input params: (array, element)
*/
func ArrayIndexOf() {
	intArr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original int array:", intArr)

	index := gouse.IndexOfArr(intArr, 3)
	fmt.Println("Index of 3:", index)

	strArr := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Println("\nOriginal string array:", strArr)

	index = gouse.IndexOfArr(strArr, "cherry")
	fmt.Println("Index of 'cherry':", index)

	println("[int]: ", gouse.IndexOfArr([]int{1, -2, 3, -4, 5, 6}, 3))
	println("[uint]: ", gouse.IndexOfArr([]uint{1, 2, 3, 4, 5, 7}, 3))
	println("[float]: ", gouse.IndexOfArr([]float64{1.2, 2.3, 3.4}, 3.4))
	println("[string]: ", gouse.IndexOfArr([]string{"1", "4", "5", "6"}, "5"))
	println("[rune]: ", gouse.IndexOfArr([]rune{'a', 'b', 'd', 'e', 'f'}, 'e'))
	println("[complex]: ", gouse.IndexOfArr([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, 5+6i))
	println("[struct]: ", gouse.IndexOfArr([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, struct{ a int }{3}))
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
Description: Removing falsy values (false, null, 0, "", undefined, and NaN) from an array
Input params: (array)
*/
func ArrayCompact() {
	result := gouse.Compact([]interface{}{1, -2, 3, -4, 5, 6, 0, 0.0, "", false, nil})
	gouse.Println("Compact remove all falsy values: ", result)
}

/*
Description: Sort array
Input params: (array)
*/
func ArraySort() {
	intArray := []int{5, 2, 9, 1, 5, 6}
	gouse.Sort(intArray)
	fmt.Println("Sorted integers:", intArray)

	strArray := []string{"apple", "orange", "banana", "grape"}
	gouse.Sort(strArray)
	fmt.Println("Sorted strings:", strArray)

	floatArray := []float64{5.2, 3.1, 7.4, 1.9}
	gouse.Sort(floatArray)
	fmt.Println("Sorted floats:", floatArray)
}

/*
Description: Flatten array
Input params: (array)
*/
func ArrayFlatten() {
	fmt.Println(gouse.Flatten([]interface{}{1, 2, 3, []interface{}{4, 5, 6, []interface{}{7, 8, 9}}}))
}

/*
Description: Shuffle array
Input params: (array)
*/
func ArrayShuffle() {
	intArr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Original int Array:", intArr)
	gouse.Shuffle(intArr)
	fmt.Println("Shuffled int Array:", intArr)

	stringArr := []string{"apple", "banana", "cherry", "date"}
	fmt.Println("Original string Array:", stringArr)
	gouse.Shuffle(stringArr)
	fmt.Println("Shuffled string Array:", stringArr)

	byteArr := []byte{'a', 'b', 'c', 'd', 'e'}
	fmt.Println("Original byte Array:", byteArr)
	gouse.Shuffle(byteArr)
	fmt.Println("Shuffled byte Array:", byteArr)
}

func ArrayIsEqual() {
	intArr1 := []int{1, 2, 3, 4}
	intArr2 := []int{1, 2, 3, 4}
	intArr3 := []int{4, 3, 2, 1}

	fmt.Println("intArr1 == intArr2:", gouse.IsEqualArr(intArr1, intArr2))
	fmt.Println("intArr1 == intArr3:", gouse.IsEqualArr(intArr1, intArr3))
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
