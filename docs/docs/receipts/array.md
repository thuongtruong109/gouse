
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Array' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Array min

Description: Find min element in array<br>Input params: (array)<br>

```go
func ArrayMin() {
	println("[int]: ", gouse.MinArr([]int{1, -2, 3}))
	println("[uint]: ", gouse.MinArr([]uint{1, 2, 3}))
	println("[string]: ", gouse.MinArr([]string{"z", "d", "m"}))
	println("[rune]: ", string(gouse.MinArr([]rune{'z', 'd', 'm'})))
	println("[float]: ", gouse.MinArr([]float64{1.2, 2.3, 3.4}))
}
```

## 2. Array max

Description: Find max element in array<br>Input params: (array)<br>

```go
func ArrayMax() {
	println("[int]: ", gouse.MaxArr([]int{1, -2, 3}))
	println("[uint]: ", gouse.MaxArr([]uint{1, 2, 3}))
	println("[string]: ", gouse.MaxArr([]string{"z", "d", "m"}))
	println("[rune]: ", string(gouse.MaxArr([]rune{'z', 'd', 'm'})))
	println("[float]: ", gouse.MaxArr([]float64{1.2, 2.3, 3.4}))
}
```

## 3. Array intersect

Description: Find intersection elements among arrays<br>Input params: (array1, array2, array3, ...)<br>

```go
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
```

## 4. Array includes

Description: Check element is exist in array<br>Input params: (array, element)<br>

```go
func ArrayIncludes() {
	println("[int]: ", gouse.IncludesArr([]int{1, -2, 3}, 1))
	println("[uint]: ", gouse.IncludesArr([]uint{1, 2, 3}, 1))
	println("[float]: ", gouse.IncludesArr([]float64{1.2, 2.3, 3.4}, 1.2))
	println("[string]: ", gouse.IncludesArr([]string{"1", "2", "3"}, "0"))
	println("[rune]: ", gouse.IncludesArr([]rune{'a', 'b', 'c'}, 'a'))
	println("[complex]: ", gouse.IncludesArr([]complex128{1 + 2i, 2 + 3i}, 1+2i))
	println("[struct]: ", gouse.IncludesArr([]struct{ a int }{{1}, {2}}, struct{ a int }{3}))
}
```

## 5. Array equal

Description: Compare is equal between two elements<br>Input params: (element1, element2)<br>

```go
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
```

## 6. Array unique

Description: Find unique elements in array<br>Input params: (array)<br>

```go
func ArrayUnique() {
	println("[int]: ", gouse.UniqueArr([]int{1, -2, 3, 1, -2, 3}))
	println("[uint]: ", gouse.UniqueArr([]uint{1, 2, 3, 1, 2, 3}))
	println("[float]: ", gouse.UniqueArr([]float64{1.2, 2.3, 3.4, 1.2, 2.3, 3.4}))
	println("[string]: ", gouse.UniqueArr([]string{"1", "4", "5", "1", "4", "5"}))
	println("[rune]: ", gouse.UniqueArr([]rune{'a', 'b', 'd', 'a', 'b', 'd'}))
}
```

## 7. Array most

Description: Find most frequency element in array<br>Input params: (array)<br>

```go
func ArrayMost() {
	println("[int]: ", gouse.Most([]int{1, -2, 3, 2, 2, 1, 2, 3}))
	println("[uint]: ", gouse.Most([]uint{1, 2, 3, 2, 2, 1, 2, 3}))
	gouse.Println("[float]: ", gouse.Most([]float64{1.2, 2.3, 3.4, 2.3, 2.3, 1.2, 2.3, 3.4}))
	println("[string]: ", gouse.Most([]string{"1", "2", "3", "2", "2", "1", "2", "3"}))
	println("[rune]: ", string(gouse.Most([]rune{'a', 'b', 'c', 'b', 'b', 'a', 'b', 'c'})))
	gouse.Println("[complex]: ", gouse.Most([]complex128{1 + 2i, 2 + 3i, 3 + 4i, 2 + 3i, 2 + 3i, 1 + 2i, 2 + 3i, 3 + 4i}))
	gouse.Println("[struct]: ", gouse.Most([]struct{ a int }{{1}, {2}, {3}, {2}, {2}, {1}, {2}, {3}}))
}
```

## 8. Array least

Description: Find least frequency element in array<br>Input params: (array)<br>

```go
func ArrayLeast() {
	println("[int]: ", gouse.Least([]int{1, -2, 3, 2, 2, 1, 2, 3}))
	println("[uint]: ", gouse.Least([]uint{1, 2, 3, 2, 2, 1, 2, 3}))
	gouse.Println("[float]: ", gouse.Least([]float64{1.2, 2.3, 3.4, 2.3, 2.3, 1.2, 2.3, 3.4}))
	println("[string]: ", gouse.Least([]string{"1", "2", "3", "2", "2", "1", "2", "3"}))
	println("[rune]: ", string(gouse.Least([]rune{'a', 'b', 'c', 'b', 'b', 'a', 'b', 'c'})))
	gouse.Println("[complex]: ", gouse.Least([]complex128{1 + 2i, 2 + 3i, 3 + 4i, 2 + 3i, 2 + 3i, 1 + 2i, 2 + 3i, 3 + 4i}))
	gouse.Println("[struct]: ", gouse.Least([]struct{ a int }{{1}, {2}, {3}, {2}, {2}, {1}, {2}, {3}}))
}
```

## 9. Array chunk

Description: Chunk an array into smaller arrays of a specified size<br>Input params: (array, size)<br>

```go
func ArrayChunk() {
	gouse.Println("[int]: ", gouse.Chunk([]int{1, -2, 3, -4, 5, 6}, 3))
	gouse.Println("[uint]: ", gouse.Chunk([]uint{1, 2, 3, 4, 5, 6}, 3))
	gouse.Println("[float]: ", gouse.Chunk([]float64{1.2, 2.3, 3.4, 4.5, 5.6, 6.7}, 3))
	gouse.Println("[string]: ", gouse.Chunk([]string{"1", "2", "3", "4", "5", "6"}, 3))
	gouse.Println("[rune]: ", gouse.Chunk([]rune{'a', 'b', 'c', 'd', 'e', 'f'}, 3))
	gouse.Println("[complex]: ", gouse.Chunk([]complex128{1 + 2i, 2 + 3i, 3 + 4i, 4 + 5i, 5 + 6i, 6 + 7i}, 3))
	gouse.Println("[struct]: ", gouse.Chunk([]struct{ a int }{{1}, {2}, {3}, {4}, {5}, {6}}, 3))
}
```

## 10. Array difference

Description: Check difference items between two arrays<br>Input params: (array1, array2)<br>

```go
func ArrayDifference() {
	gouse.Println("[int]: ", gouse.Diff([]int{1, -2, 3, -4, 5, 6}, []int{1, 2, 3, 4, 5, 6}))
	gouse.Println("[uint]: ", gouse.Diff([]uint{1, 2, 3, 4, 5, 7}, []uint{1, 2, 3, 4, 5, 6}))
	gouse.Println("[float]: ", gouse.Diff([]float64{1.2, 2.3, 3.4}, []float64{4.5, 5.6, 6.7}))
	gouse.Println("[string]: ", gouse.Diff([]string{"1", "4", "5", "6"}, []string{"1", "2", "3", "6"}))
	gouse.Println("[rune]: ", gouse.Diff([]rune{'a', 'b', 'd', 'e', 'f'}, []rune{'a', 'b', 'c', 'f'}))
	gouse.Println("[complex]: ", gouse.Diff([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, []complex128{1 + 2i, 2 + 3i, 6 + 7i}))
	gouse.Println("[struct]: ", gouse.Diff([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, []struct{ a int }{{1}, {2}}))
}
```

## 11. Array drop

Description: Drop n elements in array (default n = 1)<br>Input params: (array, n)<br>

```go
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
```

## 12. Array fill

Description: Fill array with value from start to end index<br>Input params: (array, value, start, end)<br>

```go
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
```

## 13. Array shift

Description: Shift array to left<br>Input params: (array)<br>

```go
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
```

## 14. Array unshift

Description: Unshift array to right<br>Input params: (array, values)<br>

```go
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
```

## 15. Array pop

Description: Pop n elements in array (default n = 1)<br>Input params: (array, n)<br>

```go
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
```

## 16. Array push

Description: Push element to array<br>Input params: (array, values)<br>

```go
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
```

## 17. Array slice

Description: Get elements from start to end index<br>Input params: (array, start, end)<br>

```go
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
```

## 18. Array splice

Description: Remove items from start index with n elements<br>Input params: (array, start, n)<br>

```go
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
```

## 19. Array take

Description: Take n elements from array<br>Input params: (array, n)<br>

```go
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
```

## 20. Array reverse

Description: Reverse array<br>Input params: (array)<br>

```go
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
```

## 21. Array index of

Description: Find index of element in array<br>Input params: (array, element)<br>

```go
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
```

## 22. Array merge

Description: Merge arrays<br>Input params: (array1, array2, array3, ...)<br>

```go
func ArrayMerge() {
	println("[int]: ", gouse.Merge([]int{1, -2, 3, -4, 5, 6}, []int{1, 2, 3, 4, 5, 6}, []int{1, -2, 3, -4, 5, 6}))
	println("[uint]: ", gouse.Merge([]uint{1, 2, 3, 4, 5, 7}, []uint{1, 2, 3, 4, 5, 6}))
	println("[float]: ", gouse.Merge([]float64{1.2, 2.3, 3.4}, []float64{1.2, 4.5, 5.6, 6.7}))
	println("[string]: ", gouse.Merge([]string{"1", "4", "5", "6"}, []string{"1", "2", "3", "6"}))
	println("[rune]: ", gouse.Merge([]rune{'a', 'b', 'd', 'e', 'f'}, []rune{'a', 'b', 'c', 'f'}))
	println("[complex]: ", gouse.Merge([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, []complex128{1 + 2i, 2 + 3i, 6 + 7i}))
	println("[struct]: ", gouse.Merge([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, []struct{ a int }{{1}, {2}}))
}
```

## 23. Array compact

Description: Removing falsy values (false, null, 0, "", undefined, and NaN) from an array<br>Input params: (array)<br>

```go
func ArrayCompact() {
	result := gouse.Compact([]interface{}{1, -2, 3, -4, 5, 6, 0, 0.0, "", false, nil})
	gouse.Println("Compact remove all falsy values: ", result)
}
```

## 24. Array sort

Description: Sort array<br>Input params: (array)<br>

```go
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
```

## 25. Array flatten

Description: Flatten array<br>Input params: (array)<br>

```go
func ArrayFlatten() {
	fmt.Println(gouse.Flatten([]interface{}{1, 2, 3, []interface{}{4, 5, 6, []interface{}{7, 8, 9}}}))
}
```

## 26. Array shuffle

Description: Shuffle array<br>Input params: (array)<br>

```go
func ArrayShuffle() {
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Original int array:", intArr)

	shuffledIntArr := gouse.Shuffle(intArr)
	fmt.Println("Shuffled int array:", shuffledIntArr)

	strArr := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Println("\nOriginal string array:", strArr)

	shuffledStrArr := gouse.Shuffle(strArr)
	fmt.Println("Shuffled string array:", shuffledStrArr)
}
```

## 27. Array filter by

Description: Filter elements in array by pass condition in callback function<br>Input params: (array, callback)<br>

```go
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
```

## 28. Array find by

Description: Find element in array by pass condition in callback function<br>Input params: (array, callback)<br>

```go
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
```

## 29. Array for by

Description: Loop array then handler with callback function<br>Input params: (array, callback)<br>

```go
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
```

## 30. Array index by

Description: Find index of element pass condition in callback function<br>Input params: (array, callback)<br>

```go
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
```

## 31. Array key by

Description: Find key of element pass condition in callback function<br>Input params: (array, callback)<br>

```go
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
```

## 32. Array map by

Description: Map array then handler with callback function<br>Input params: (array, callback)<br>

```go
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
```

## 33. Array reject by

Description: Remove element in array by pass condition in callback function<br>Input params: (array, callback)<br>

```go
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
```
