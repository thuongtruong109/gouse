package gouse

import "testing"

func TestMinMaxArr(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	minExpected := 1
	if MinArr(arr) != minExpected {
		t.Errorf("Expected %d, got %d", minExpected, MinArr(arr))
	}

	maxExpected := 5
	if MaxArr(arr) != maxExpected {
		t.Errorf("Expected %d, got %d", maxExpected, MaxArr(arr))
	}
}

func TestIntersect(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		slices [][]int
		want   []int
	}{
		{
			name:   "empty slices",
			slices: [][]int{},
			want:   nil,
		},
		{
			name:   "one empty slice",
			slices: [][]int{{}, {1, 2, 3}},
			want:   nil,
		},
		{
			name:   "two slices",
			slices: [][]int{{1, 2, 3}, {2, 3, 4}},
			want:   []int{2, 3},
		},
		{
			name:   "three slices",
			slices: [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
			want:   []int{3},
		},
		{
			name:   "four slices",
			slices: [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}, {4, 5, 6}},
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Intersect(tt.slices...)

			if len(got) != len(tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Intersect() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestIndexBy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		f    func(int) bool
		want int
	}{
		{
			name: "empty array",
			arr:  []int{},
			f:    func(int) bool { return true },
			want: -1,
		},
		{
			name: "not found",
			arr:  []int{1, 2, 3},
			f:    func(v int) bool { return v == 4 },
			want: -1,
		},
		{
			name: "found",
			arr:  []int{1, 2, 3},
			f:    func(v int) bool { return v == 2 },
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IndexBy(tt.arr, tt.f)

			if got != tt.want {
				t.Errorf("IndexBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyBy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		f    func(int) bool
		want int
	}{
		{
			name: "empty array",
			arr:  []int{},
			f:    func(int) bool { return true },
			want: -1,
		},
		{
			name: "not found",
			arr:  []int{1, 2, 3},
			f:    func(v int) bool { return v == 4 },
			want: -1,
		},
		{
			name: "found",
			arr:  []int{1, 2, 3},
			f:    func(v int) bool { return v == 2 },
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KeyBy(tt.arr, tt.f)

			if got != tt.want {
				t.Errorf("KeyBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterBy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		f    func(int) bool
		want []int
	}{
		{
			name: "empty array",
			arr:  []int{},
			f:    func(int) bool { return true },
			want: nil,
		},
		{
			name: "not found",
			arr:  []int{1, 2, 3},
			f:    func(v int) bool { return v == 4 },
			want: nil,
		},
		{
			name: "found",
			arr:  []int{1, 2, 3},
			f:    func(v int) bool { return v == 2 },
			want: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterBy(tt.arr, tt.f)

			if len(got) != len(tt.want) {
				t.Errorf("FilterBy() = %v, want %v", got, tt.want)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("FilterBy() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestRejectBy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		f    func(int) bool
		want []int
	}{
		{
			name: "empty array",
			arr:  []int{},
			f:    func(int) bool { return true },
			want: nil,
		},
		{
			name: "not found",
			arr:  []int{1, 2, 3},
			f:    func(v int) bool { return v == 4 },
			want: []int{1, 2, 3},
		},
		{
			name: "found",
			arr:  []int{1, 2, 3},
			f:    func(v int) bool { return v == 2 },
			want: []int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RejectBy(tt.arr, tt.f)

			if len(got) != len(tt.want) {
				t.Errorf("RejectBy() = %v, want %v", got, tt.want)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("RejectBy() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestFindBy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		f    func(int) bool
		want int
	}{
		{
			name: "not found",
			arr:  []int{1, 2, 3},
			f:    func(v int) bool { return v == 4 },
			want: 1,
		},
		{
			name: "found",
			arr:  []int{1, 2, 3},
			f:    func(v int) bool { return v == 2 },
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindBy(tt.arr, tt.f)

			if got != tt.want {
				t.Errorf("FindBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForBy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		f    func(int)
	}{
		{
			name: "empty array",
			arr:  []int{},
			f:    func(int) {},
		},
		{
			name: "not found",
			arr:  []int{1, 2, 3},
			f:    func(int) {},
		},
		{
			name: "found",
			arr:  []int{1, 2, 3},
			f:    func(int) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ForBy(tt.arr, tt.f)
		})
	}
}

func TestMapBy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		f    func(int) int
		want []int
	}{
		{
			name: "empty array",
			arr:  []int{},
			f:    func(int) int { return 0 },
			want: nil,
		},
		{
			name: "not found",
			arr:  []int{1, 2, 3},
			f:    func(v int) int { return v + 1 },
			want: []int{2, 3, 4},
		},
		{
			name: "found",
			arr:  []int{1, 2, 3},
			f:    func(v int) int { return v * 2 },
			want: []int{2, 4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapBy(tt.arr, tt.f)

			if len(got) != len(tt.want) {
				t.Errorf("MapBy() = %v, want %v", got, tt.want)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("MapBy() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestIncludesArr(t *testing.T) {
	if !IncludesArr([]int{1, 2, 3}, 1) {
		t.Errorf("Exist[int]([]int{1, 2, 3}, 1) = false, want true")
	}
}

func TestEqual(t *testing.T) {
	if !IncludesArr([]int{1, 2, 3}, 1) {
		t.Errorf("Equal[int](1, 1) = false, want true")
	}
}

func TestSumArr(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	expected := 15
	if SumArr(arr) != expected {
		t.Errorf("Expected %d, got %d", expected, SumArr(arr))
	}
}

func TestProductArr(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	expected := 120
	if ProductArr(arr) != expected {
		t.Errorf("Expected %d, got %d", expected, ProductArr(arr))
	}
}

func TestUniqueArr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "empty array",
			arr:  []int{},
			want: nil,
		},
		{
			name: "no duplicates",
			arr:  []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "duplicates",
			arr:  []int{1, 2, 2, 3},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniqueArr(tt.arr)

			if len(got) != len(tt.want) {
				t.Errorf("UniqueArr() = %v, want %v", got, tt.want)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("UniqueArr() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestMost(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		want int
	}{
		// {
		// 	name: "empty array",
		// 	arr:  []int{},
		// 	want: 0,
		// },
		{
			name: "single element",
			arr:  []int{1},
			want: 1,
		},
		{
			name: "multiple elements",
			arr:  []int{1, 2, 2, 3},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Most(tt.arr)

			if got != tt.want {
				t.Errorf("Most() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		size int
		want [][]int
	}{
		{
			name: "empty array",
			arr:  []int{},
			size: 2,
			want: nil,
		},
		{
			name: "single element",
			arr:  []int{1},
			size: 2,
			want: [][]int{{1}},
		},
		{
			name: "multiple elements",
			arr:  []int{1, 2, 3, 4, 5},
			size: 2,
			want: [][]int{{1, 2}, {3, 4}, {5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Chunk(tt.arr, tt.size)

			if len(got) != len(tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}

			for i := range got {
				if len(got[i]) != len(tt.want[i]) {
					t.Errorf("Chunk() = %v, want %v", got, tt.want)
				}

				for j := range got[i] {
					if got[i][j] != tt.want[i][j] {
						t.Errorf("Chunk() = %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}

func TestDiff(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{
			name: "empty arrays",
			a:    []int{},
			b:    []int{},
			want: nil,
		},
		{
			name: "no difference",
			a:    []int{1, 2, 3},
			b:    []int{1, 2, 3},
			want: nil,
		},
		{
			name: "difference",
			a:    []int{1, 2, 3},
			b:    []int{2, 3, 4},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Diff(tt.a, tt.b)

			if len(got) != len(tt.want) {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Diff() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestDrop(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		n    int
		want []int
	}{
		{
			name: "empty array",
			arr:  []int{},
			n:    1,
			want: nil,
		},
		{
			name: "no drop",
			arr:  []int{1, 2, 3},
			n:    0,
			want: []int{1, 2, 3},
		},
		{
			name: "drop one",
			arr:  []int{1, 2, 3},
			n:    1,
			want: []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Drop(tt.arr, tt.n)

			if len(got) != len(tt.want) {
				t.Errorf("Drop() = %v, want %v", got, tt.want)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Drop() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestIndexOfArr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		arr   []int
		value int
		want  int
	}{
		{
			name:  "empty array",
			arr:   []int{},
			value: 1,
			want:  -1,
		},
		{
			name:  "not found",
			arr:   []int{1, 2, 3},
			value: 4,
			want:  -1,
		},
		{
			name:  "found",
			arr:   []int{1, 2, 3},
			value: 2,
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IndexOfArr(tt.arr, tt.value)

			if got != tt.want {
				t.Errorf("IndexOfArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  [][]int
		want []int
	}{
		{
			name: "empty array",
			arr:  [][]int{},
			want: nil,
		},
		{
			name: "no merge",
			arr:  [][]int{{1}, {2}, {3}},
			want: []int{1, 2, 3},
		},
		{
			name: "merge",
			arr:  [][]int{{1}, {2}, {3}, {4}},
			want: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.arr...)

			if len(got) != len(tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Merge() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestCompact(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "empty array",
			arr:  []int{},
			want: nil,
		},
		{
			name: "no compact",
			arr:  []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "compact",
			arr:  []int{0, 1, 2, 3},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Compact(tt.arr)

			if len(got) != len(tt.want) {
				t.Errorf("Compact() = %v, want %v", got, tt.want)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Compact() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

// func TestSort(t *testing.T) {
// 	t.Parallel()

// 	tests := []struct {
// 		name string
// 		arr  []int
// 		want []int
// 	}{
// 		{
// 			name: "empty array",
// 			arr:  []int{},
// 			want: nil,
// 		},
// 		{
// 			name: "no sort",
// 			arr:  []int{3, 2, 1},
// 			want: []int{1, 2, 3},
// 		},
// 		{
// 			name: "sort",
// 			arr:  []int{3, 1, 2},
// 			want: []int{1, 2, 3},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			Sort(tt.arr)

// 			for i := range tt.arr {
// 				if tt.arr[i] != tt.want[i] {
// 					t.Errorf("Sort() = %v, want %v", tt.arr, tt.want)
// 				}
// 			}
// 		})
// 	}
// }
