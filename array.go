package gouse

import (
	"math/rand"
	"reflect"
	"sort"
	"time"
)

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

func MinArr[T Number | string](arr []T) T {
	return _minmax(arr, func(a, b T) bool {
		return a > b
	})
}

func MaxArr[T Number | string](arr []T) T {
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

func IncludesArr[T comparable](array []T, value T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
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

func Least[T comparable](arr []T) T {
	var least = make(map[T]int)
	min := arr[0]
	for _, v := range arr {
		least[v] = least[v] + 1
		if least[min] > least[v] {
			min = v
		}
	}

	return min
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

func Fill[T comparable](arr []T, value T, start, end int) []T {
	if len(arr) == 0 {
		return nil
	}

	if start < 0 {
		start = 0
	}

	if end > len(arr) {
		end = len(arr)
	}

	for i := start; i < end; i++ {
		arr[i] = value
	}

	return arr
}

func Shift[T comparable](arr []T, n ...int) []T {
	if len(arr) == 0 {
		return nil
	}

	if len(n) == 0 {
		n = append(n, 1)
	}

	return arr[:len(arr)-n[0]]
}

func Unshift[T comparable](arr []T, values ...T) []T {
	return append(values, arr...)
}

func Pop[T comparable](arr []T, n ...int) []T {
	if len(arr) == 0 {
		return nil
	}

	if len(n) == 0 {
		n = append(n, 1)
	}

	if n[0] >= len(arr) {
		return []T{}
	}

	return arr[:len(arr)-n[0]]
}

func Push[T comparable](arr []T, values ...T) []T {
	return append(arr, values...)
}

func SliceArr[T comparable](arr []T, start, end int) ([]T, error) {
	if len(arr) == 0 {
		return nil, nil
	}

	if start < 0 {
		start = 0
	}

	if end > len(arr) {
		end = len(arr)
	}

	if start >= len(arr) {
		return nil, nil
	}

	return arr[start:end], nil
}

func SpliceArr[T comparable](arr []T, start, deleteCount int, items ...T) []T {
	if len(arr) == 0 && start == 0 && deleteCount == 0 {
		return items
	}

	if start < 0 {
		start = 0
	}

	if start > len(arr) {
		start = len(arr)
	}

	if deleteCount < 0 {
		deleteCount = 0
	}
	if deleteCount > len(arr)-start {
		deleteCount = len(arr) - start
	}

	spliced := append(arr[:start], items...)

	spliced = append(spliced, arr[start+deleteCount:]...)

	return spliced
}

func Take[T comparable](arr []T, n int) []T {
	if len(arr) == 0 {
		return nil
	}

	if n > len(arr) {
		n = len(arr)
	}

	return arr[:n]
}

func ReverseArr[T comparable](arr []T) []T {
	var reversed []T
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	return reversed
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

func Sort(input any) any {
	switch v := input.(type) {
	case []int:
		sort.Ints(v)
		return v
	case []string:
		sort.Strings(v)
		return v
	case []float64:
		sort.Float64s(v)
		return v
	default:
		return nil
	}
}

func Flatten(input any) []any {
	var result []any

	v := reflect.ValueOf(input)
	if v.Kind() == reflect.Slice {
		for i := range v.Len() {
			elem := v.Index(i).Interface()
			if reflect.ValueOf(elem).Kind() == reflect.Slice {
				result = append(result, Flatten(elem)...)
			} else {
				result = append(result, elem)
			}
		}
	}

	return result
}

func Shuffle[T comparable](arr []T) {
	rand.Seed(time.Now().UnixNano())

	originalArr := append([]T(nil), arr...)
	for {
		for i := len(arr) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			arr[i], arr[j] = arr[j], arr[i]
		}

		if !IsEqualArr(arr, originalArr) {
			break
		}
	}
}

func IsEqualArr[T comparable](a, b []T) bool {
	return reflect.DeepEqual(a, b)
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
