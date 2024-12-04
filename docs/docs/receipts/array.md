
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Array' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

### <Badge style='font-size: 1.1rem;' type='tip' text='1. array chunk' />

Description: Chunk an array into smaller arrays of a specified size<br>Input params: (array, size)<br>

```go
func ArrayChunk() {
	fmt.Println("[int]: ", gouse.Chunk([]int{1, -2, 3, -4, 5, 6}, 3))
	fmt.Println("[uint]: ", gouse.Chunk([]uint{1, 2, 3, 4, 5, 6}, 3))
	fmt.Println("[float]: ", gouse.Chunk([]float64{1.2, 2.3, 3.4, 4.5, 5.6, 6.7}, 3))
	fmt.Println("[string]: ", gouse.Chunk([]string{"1", "2", "3", "4", "5", "6"}, 3))
	fmt.Println("[rune]: ", gouse.Chunk([]rune{'a', 'b', 'c', 'd', 'e', 'f'}, 3))
	fmt.Println("[complex]: ", gouse.Chunk([]complex128{1 + 2i, 2 + 3i, 3 + 4i, 4 + 5i, 5 + 6i, 6 + 7i}, 3))
	fmt.Println("[struct]: ", gouse.Chunk([]struct{ a int }{{1}, {2}, {3}, {4}, {5}, {6}}, 3))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='2. array compact' />

Description: Removing falsy values (false, null, 0, "", undefined, and NaN) from an array<br>Input params: (array)<br>

```go
func ArrayCompact() {
	result := gouse.Compact([]interface{}{1, -2, 3, -4, 5, 6, 0, 0.0, "", false, nil})
	fmt.Println("Compact remove all falsy values: ", result)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='3. array difference' />

Description: Check difference items between two arrays<br>Input params: (array1, array2)<br>

```go
func ArrayDifference() {
	fmt.Println("[int]: ", gouse.Diff([]int{1, -2, 3, -4, 5, 6}, []int{1, 2, 3, 4, 5, 6}))
	fmt.Println("[uint]: ", gouse.Diff([]uint{1, 2, 3, 4, 5, 7}, []uint{1, 2, 3, 4, 5, 6}))
	fmt.Println("[float]: ", gouse.Diff([]float64{1.2, 2.3, 3.4}, []float64{4.5, 5.6, 6.7}))
	fmt.Println("[string]: ", gouse.Diff([]string{"1", "4", "5", "6"}, []string{"1", "2", "3", "6"}))
	fmt.Println("[rune]: ", gouse.Diff([]rune{'a', 'b', 'd', 'e', 'f'}, []rune{'a', 'b', 'c', 'f'}))
	fmt.Println("[complex]: ", gouse.Diff([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, []complex128{1 + 2i, 2 + 3i, 6 + 7i}))
	fmt.Println("[struct]: ", gouse.Diff([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, []struct{ a int }{{1}, {2}}))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='4. array drop' />

Description: Drop n elements in array (default n = 1)<br>Input params: (array, n)<br>

```go
func ArrayDrop() {
	fmt.Println("[int] with default: ", gouse.Drop([]int{1, -2, 3, -4, 5, 6}))
	fmt.Println("[int]: ", gouse.Drop([]int{1, -2, 3, -4, 5, 6}, 2))
	fmt.Println("[uint]: ", gouse.Drop([]uint{1, 2, 3, 4, 5, 7}, 2))
	fmt.Println("[float]: ", gouse.Drop([]float64{1.2, 2.3, 3.4}, 2))
	fmt.Println("[string]: ", gouse.Drop([]string{"1", "4", "5", "6"}, 2))
	fmt.Println("[rune]: ", gouse.Drop([]rune{'a', 'b', 'd', 'e', 'f'}, 2))
	fmt.Println("[complex]: ", gouse.Drop([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, 2))
	fmt.Println("[struct]: ", gouse.Drop([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, 2))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='5. array equal' />

Description: Compare is equal between two elements<br>Input params: (element1, element2)<br>

```go
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
```

### <Badge style='font-size: 1.1rem;' type='tip' text='6. array filter by' />

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

### <Badge style='font-size: 1.1rem;' type='tip' text='7. array find by' />

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

	fmt.Println("[struct]: ", gouse.FindBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) bool {
		return v.a == 3
	}))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='8. array for by' />

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
		fmt.Println(v)
	})
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='9. array includes' />

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

### <Badge style='font-size: 1.1rem;' type='tip' text='10. array index' />

Description: Get index of element in array<br>Input params: (array, element)<br>

```go
func ArrayIndex() {
	println("[int]: ", gouse.IndexOfArr([]int{1, -2, 3, -4, 5, 6}, 3))
	println("[uint]: ", gouse.IndexOfArr([]uint{1, 2, 3, 4, 5, 7}, 3))
	println("[float]: ", gouse.IndexOfArr([]float64{1.2, 2.3, 3.4}, 3.4))
	println("[string]: ", gouse.IndexOfArr([]string{"1", "4", "5", "6"}, "5"))
	println("[rune]: ", gouse.IndexOfArr([]rune{'a', 'b', 'd', 'e', 'f'}, 'e'))
	println("[complex]: ", gouse.IndexOfArr([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, 5+6i))
	println("[struct]: ", gouse.IndexOfArr([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, struct{ a int }{3}))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='11. array index by' />

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

### <Badge style='font-size: 1.1rem;' type='tip' text='12. array intersect' />

Description: Intersection arrays<br>Input params: (array1, array2)<br>

```go
func ArrayIntersect() {
	println("[int]: ", gouse.Intersect([]int{1, -2, 3, -4, 5, 6}, []int{1, 2, 3, 4, 5, 6}))
	println("[uint]: ", gouse.Intersect([]uint{1, 2, 3, 4, 5, 7}, []uint{1, 2, 3, 4, 5, 6}))
	println("[float]: ", gouse.Intersect([]float64{1.2, 2.3, 3.4}, []float64{1.2, 4.5, 5.6, 6.7}))
	println("[string]: ", gouse.Intersect([]string{"1", "4", "5", "6"}, []string{"1", "2", "3", "6"}))
	println("[rune]: ", gouse.Intersect([]rune{'a', 'b', 'd', 'e', 'f'}, []rune{'a', 'b', 'c', 'f'}))
	println("[complex]: ", gouse.Intersect([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, []complex128{1 + 2i, 2 + 3i, 6 + 7i}))
	println("[struct]: ", gouse.Intersect([]struct{ a int }{{1}, {-2}, {3}, {4}, {5}, {6}}, []struct{ a int }{{1}, {2}}))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='13. array key by' />

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

### <Badge style='font-size: 1.1rem;' type='tip' text='14. array map by' />

Description: Map array then handler with callback function<br>Input params: (array, callback)<br>

```go
func ArrayMapBy() {
	fmt.Println("[int]: ", gouse.MapBy([]int{1, -2, 3, -4, 5, 6}, func(v int) int {
		return v * 2
	}))

	fmt.Println("[uint]: ", gouse.MapBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) uint {
		return v * 2
	}))

	fmt.Println("[float]: ", gouse.MapBy([]float64{1.2, 2.3, 3.4}, func(v float64) float64 {
		return v * 2
	}))

	fmt.Println("[string]: ", gouse.MapBy([]string{"1", "4", "5", "6"}, func(v string) string {
		return v + "1"
	}))

	fmt.Println("[rune]: ", gouse.MapBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) rune {
		return v + 1
	}))

	fmt.Println("[complex]: ", gouse.MapBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) complex128 {
		return v * 2
	}))

	fmt.Println("[struct]: ", gouse.MapBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) struct{ a int } {
		return struct{ a int }{v.a * 2}
	}))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='15. array min' />

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

### <Badge style='font-size: 1.1rem;' type='tip' text='16. array max' />

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

### <Badge style='font-size: 1.1rem;' type='tip' text='17. array most' />

Description: Find most frequency element in array<br>Input params: (array)<br>

```go
func ArrayMost() {
	println("[int]: ", gouse.Most([]int{1, -2, 3, 2, 2, 1, 2, 3}))
	println("[uint]: ", gouse.Most([]uint{1, 2, 3, 2, 2, 1, 2, 3}))
	fmt.Println("[float]: ", gouse.Most([]float64{1.2, 2.3, 3.4, 2.3, 2.3, 1.2, 2.3, 3.4}))
	println("[string]: ", gouse.Most([]string{"1", "2", "3", "2", "2", "1", "2", "3"}))
	println("[rune]: ", string(gouse.Most([]rune{'a', 'b', 'c', 'b', 'b', 'a', 'b', 'c'})))
	fmt.Println("[complex]: ", gouse.Most([]complex128{1 + 2i, 2 + 3i, 3 + 4i, 2 + 3i, 2 + 3i, 1 + 2i, 2 + 3i, 3 + 4i}))
	fmt.Println("[struct]: ", gouse.Most([]struct{ a int }{{1}, {2}, {3}, {2}, {2}, {1}, {2}, {3}}))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='18. array merge' />

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

### <Badge style='font-size: 1.1rem;' type='tip' text='19. array reject by' />

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

### <Badge style='font-size: 1.1rem;' type='tip' text='20. array sum' />

Description: Calculate sum of elements in array<br>Input params: (array)<br>

```go
func ArraySum() {
	println("[int]: ", gouse.SumArr([]int{1, -2, 3}))
	println("[uint]: ", gouse.SumArr([]uint{1, 2, 3}))
	fmt.Println("[float]: ", gouse.SumArr([]float64{1.2, 2.3, 3.4}))
	println("[rune]: ", gouse.SumArr([]rune{'a', 'b', 'c'}))
	fmt.Println("[complex]: ", gouse.SumArr([]complex128{1 + 2i, 2 + 3i, 3 + 4i}))
}
```
