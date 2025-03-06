package gouse

func _minmax[T comparable](arr []T, less func(T, T) bool) T {
	if len(arr) == 0 {
		panic("Empty array")
	}

	max := arr[0]
	for _, item := range arr {
		if less(max, item) {
			max = item
		}
	}
	return max
}

func MinArr[T int | int8 | int16 | int32 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string](arr []T) T {
	return _minmax(arr, func(a, b T) bool {
		return a > b
	})
}

func MaxArr[T int | int8 | int16 | int32 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string](arr []T) T {
	return _minmax(arr, func(a, b T) bool {
		return a < b
	})
}

func _interSlice[T comparable](a, b []T) []T {
	var intersect []T

	for _, v := range a {
		if IncludesArr(b, v) {
			intersect = append(intersect, v)
		}
	}

	return intersect
}

func Intersect[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}

	intersect := slices[0]

	for _, slice := range slices[1:] {
		intersect = _interSlice(intersect, slice)
	}

	return intersect
}

func IndexBy[T comparable](arr []T, f func(T) bool) int {
	for i, v := range arr {
		if f(v) {
			return i
		}
	}
	return -1
}

func KeyBy[T comparable](arr []T, f func(T) bool) int {
	return IndexBy(arr, f)
}

func FilterBy[T comparable](arr []T, f func(T) bool) []T {
	var res []T
	for _, v := range arr {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func RejectBy[T comparable](arr []T, f func(T) bool) []T {
	var res []T
	for _, v := range arr {
		if !f(v) {
			res = append(res, v)
		}
	}
	return res
}

func FindBy[T comparable](arr []T, f func(T) bool) T {
	for _, v := range arr {
		if f(v) {
			return v
		}
	}
	return arr[0]
}

func ForBy[T comparable](arr []T, f func(T)) {
	for _, v := range arr {
		f(v)
	}
}

func MapBy[T comparable, R comparable](arr []T, f func(T) R) []R {
	var res []R
	for _, v := range arr {
		res = append(res, f(v))
	}
	return res
}

func IncludesArr[T comparable](array []T, value T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func Equal[T comparable](a, b T) bool {
	return a == b
}

func SumArr[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128](arr []T) T {
	var sum T
	for _, v := range arr {
		sum += v
	}
	return sum
}

func ProductArr[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](arr []T) T {
	var product T = 1
	for _, v := range arr {
		product *= v
	}
	return product
}

func UniqueArr[T comparable](arr []T) []T {
	var unique []T
	for _, v := range arr {
		if !IncludesArr(unique, v) {
			unique = append(unique, v)
		}
	}
	return unique
}

func Most[T comparable](arr []T) T {
	var most = make(map[T]int)
	max := arr[0]
	for _, v := range arr {
		most[v] = most[v] + 1
		if most[max] < most[v] {
			max = v
		}
	}

	return max
}

func Chunk[T comparable](array []T, size int) [][]T {
	var chunks [][]T
	for i := 0; i < len(array); i += size {
		end := i + size
		if end > len(array) {
			end = len(array)
		}
		chunks = append(chunks, array[i:end])
	}
	return chunks
}

func Diff[T comparable](a, b []T) []T {
	var diff []T
	for _, v := range a {
		if !IncludesArr(b, v) {
			diff = append(diff, v)
		}
	}
	return diff
}

func Drop[T comparable](arr []T, n ...int) []T {
	if len(arr) == 0 {
		return nil
	}

	if len(n) == 0 {
		n = append(n, 1)
	}

	return arr[n[0]:]
}

func IndexOfArr[T comparable](arr []T, value T) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

func Merge[T comparable](arr ...[]T) []T {
	var merged []T
	for _, v := range arr {
		merged = append(merged, v...)
	}
	return merged
}

func Compact[T any](arr []T) []T {
	var compact []T
	for _, v := range arr {
		if !IsZero(v) && !IsNil(v) && !IsEmpty(v) && !IsUndefined(v) && !IsBool(v) {
			compact = append(compact, v)
		}
	}
	return compact
}

// func Sort[T any](arr []T) []T {
// 	var sorted []T
// 	for _, v := range arr {
// 		if IsNumber(v) {
// 			sorted = append(sorted, v)
// 		}
// 	}
// 	return sorted
// }

// flatten array
