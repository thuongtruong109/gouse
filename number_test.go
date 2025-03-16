package gouse

import "testing"

func TestClamp(t *testing.T) {
	min := 1
	max := 10

	result := Clamp(0, min, max)

	if result != min {
		t.Errorf("Clamp() = %d; want %d", result, min)
	}

	result = Clamp(11, min, max)

	if result != max {
		t.Errorf("Clamp() = %d; want %d", result, max)
	}

	result = Clamp(5, min, max)

	if result != 5 {
		t.Errorf("Clamp() = %d; want 5", result)
	}
}

func BenchmarkClamp(b *testing.B) {
	min := 1
	max := 10

	for i := 0; i < b.N; i++ {
		Clamp(5, min, max)
	}
}

func TestInRange(t *testing.T) {
	min := 1
	max := 10

	result := InRange(0, min, max)

	if result {
		t.Errorf("InRange() = %t; want false", result)
	}

	result = InRange(11, min, max)

	if result {
		t.Errorf("InRange() = %t; want false", result)
	}

	result = InRange(5, min, max)

	if !result {
		t.Errorf("InRange() = %t; want true", result)
	}
}

func BenchmarkInRange(b *testing.B) {
	min := 1
	max := 10

	for i := 0; i < b.N; i++ {
		InRange(5, min, max)
	}
}

func TestSort(t *testing.T) {
	nums := []int{5, 3, 1, 4, 2}

	result := Sort(nums)

	for i := range len(result) - 1 {
		if result[i] > result[i+1] {
			t.Errorf("Sort() = %v; want [1 2 3 4 5]", result)
		}
	}

	nums2 := []float64{5.5, 3.3, 1.1, 4.4, 2.2}

	result2 := Sort(nums2)

	for i := range len(result2) - 1 {
		if result2[i] > result2[i+1] {
			t.Errorf("SortF() = %v; want [1.1 2.2 3.3 4.4 5.5]", result2)
		}
	}
}

func BenchmarkSort(b *testing.B) {
	nums := []int{5, 3, 1, 4, 2}

	for i := 0; i < b.N; i++ {
		Sort(nums)
	}

	nums2 := []float64{5.5, 3.3, 1.1, 4.4, 2.2}

	for i := 0; i < b.N; i++ {
		Sort(nums2)
	}
}
