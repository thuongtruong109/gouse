package gouse

import (
	"reflect"
	"testing"
)

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
	tests := []struct {
		a, b   []int
		expect bool
	}{
		{
			a:      []int{1, 2, 3},
			b:      []int{1, 2, 3},
			expect: true,
		},
		{
			a:      []int{1, 2, 3},
			b:      []int{1, 2, 4},
			expect: false,
		},
		{
			a:      []int{1, 2, 3},
			b:      []int{1, 2},
			expect: false,
		},
		{
			a:      []int{1, 2, 3},
			b:      []int{1, 2, 3, 4},
			expect: false,
		},
		{
			a:      []int{},
			b:      []int{},
			expect: true,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := Equal(test.a, test.b)
			if got != test.expect {
				t.Errorf("Equal(%v, %v) = %v; want %v", test.a, test.b, got, test.expect)
			}
		})
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

func TestLeast(t *testing.T) {
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
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Least(tt.arr)

			if got != tt.want {
				t.Errorf("Least() = %v, want %v", got, tt.want)
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

func TestFill(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		arr   []int
		value int
		start int
		end   int
		want  []int
	}{
		{
			name:  "empty array",
			arr:   []int{},
			value: 1,
			start: 0,
			end:   1,
			want:  nil,
		},
		{
			name:  "no fill",
			arr:   []int{1, 2, 3},
			value: 4,
			start: 0,
			end:   3,
			want:  []int{4, 4, 4},
		},
		{
			name:  "fill",
			arr:   []int{1, 2, 3},
			value: 4,
			start: 1,
			end:   2,
			want:  []int{1, 4, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Fill(tt.arr, tt.value, tt.start, tt.end)

			if len(got) != len(tt.want) {
				t.Errorf("Fill() = %v, want %v", got, tt.want)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Fill() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestShift(t *testing.T) {
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
			name: "no shift",
			arr:  []int{1, 2, 3},
			n:    0,
			want: []int{1, 2, 3},
		},
		{
			name: "shift one",
			arr:  []int{1, 2, 3},
			n:    1,
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Shift(tt.arr, tt.n)

			if len(got) != len(tt.want) {
				t.Errorf("Shift() = %v, want %v", got, tt.want)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Shift() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestUnshift(t *testing.T) {
	arr := []int{}
	result := Unshift(arr, 1)
	expected := []int{1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed, expected %v but got %v", expected, result)
	}

	arr = []int{}
	result = Unshift(arr, 1, 2, 3)
	expected = []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed, expected %v but got %v", expected, result)
	}

	arr = []int{4, 5, 6}
	result = Unshift(arr, 1, 2, 3)
	expected = []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed, expected %v but got %v", expected, result)
	}

	arr = []int{4, 5, 6}
	result = Unshift(arr)
	expected = []int{4, 5, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed, expected %v but got %v", expected, result)
	}
}

func TestPop(t *testing.T) {
	arr := []int{}
	result := Pop(arr)
	if result != nil {
		t.Errorf("Test failed, expected nil but got %v", result)
	}

	arr = []int{1, 2, 3, 4, 5}
	result = Pop(arr)
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed, expected %v but got %v", expected, result)
	}

	arr = []int{1, 2, 3, 4, 5}
	result = Pop(arr, 2)
	expected = []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed, expected %v but got %v", expected, result)
	}

	arr = []int{1, 2, 3, 4, 5}
	result = Pop(arr, 5)
	expected = []int{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed, expected %v but got %v", expected, result)
	}

	arr = []int{1, 2, 3}
	result = Pop(arr, 10)
	expected = []int{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed, expected %v but got %v", expected, result)
	}

	arr = []int{1}
	result = Pop(arr)
	expected = []int{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed, expected %v but got %v", expected, result)
	}
}

func TestPush(t *testing.T) {
	arr := []int{}
	result := Push(arr, 1)
	expected := []int{1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed, expected %v but got %v", expected, result)
	}

	arr = []int{1}
	result = Push(arr, 2, 3, 4)
	expected = []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed, expected %v but got %v", expected, result)
	}

	arr = []int{1, 2, 3}
	result = Push(arr)
	expected = []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed, expected %v but got %v", expected, result)
	}
}

func TestSliceArr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		arr   []int
		start int
		end   int
		want  []int
		err   error
	}{
		{
			name:  "empty array",
			arr:   []int{},
			start: 0,
			end:   1,
			want:  nil,
			err:   nil,
		},
		{
			name:  "no slice",
			arr:   []int{1, 2, 3},
			start: 0,
			end:   3,
			want:  []int{1, 2, 3},
			err:   nil,
		},
		{
			name:  "slice from index 1",
			arr:   []int{1, 2, 3},
			start: 1,
			end:   3,
			want:  []int{2, 3},
			err:   nil,
		},
		{
			name:  "slice with end exceeding length",
			arr:   []int{1, 2, 3},
			start: 1,
			end:   5,
			want:  []int{2, 3},
			err:   nil,
		},
		{
			name:  "slice with start exceeding length",
			arr:   []int{1, 2, 3},
			start: 5,
			end:   6,
			want:  nil,
			err:   nil,
		},
		{
			name:  "slice with negative start",
			arr:   []int{1, 2, 3},
			start: -1,
			end:   2,
			want:  []int{1, 2},
			err:   nil,
		},
		{
			name:  "slice with start equal to end",
			arr:   []int{1, 2, 3},
			start: 2,
			end:   2,
			want:  nil,
			err:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SliceArr(tt.arr, tt.start, tt.end)

			if err != nil || tt.err != nil {
				if (err == nil && tt.err != nil) || (err != nil && err.Error() != tt.err.Error()) {
					t.Errorf("SliceArr() error = %v, want %v", err, tt.err)
				}
			}

			if len(got) == 0 && tt.want != nil && len(tt.want) > 0 {
				t.Errorf("SliceArr() = %v, want %v", got, tt.want)
			} else if len(got) != len(tt.want) {
				t.Errorf("SliceArr() = %v, want %v", got, tt.want)
			} else {
				for i := range got {
					if got[i] != tt.want[i] {
						t.Errorf("SliceArr() = %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}

func TestSpliceArr(t *testing.T) {
	tests := []struct {
		name        string
		arr         []int
		start       int
		deleteCount int
		items       []int
		expected    []int
	}{
		{
			name:        "Empty array",
			arr:         []int{},
			start:       0,
			deleteCount: 0,
			items:       []int{1, 2},
			expected:    []int{1, 2},
		},
		{
			name:        "Delete elements from the middle",
			arr:         []int{1, 2, 3, 4, 5},
			start:       1,
			deleteCount: 2,
			items:       []int{8, 9},
			expected:    []int{1, 8, 9, 4, 5},
		},
		{
			name:        "Delete from the end",
			arr:         []int{1, 2, 3, 4, 5},
			start:       3,
			deleteCount: 2,
			items:       []int{},
			expected:    []int{1, 2, 3},
		},
		{
			name:        "Insert at the end",
			arr:         []int{1, 2, 3},
			start:       3,
			deleteCount: 0,
			items:       []int{4, 5},
			expected:    []int{1, 2, 3, 4, 5},
		},
		{
			name:        "Delete more elements than available",
			arr:         []int{1, 2, 3},
			start:       1,
			deleteCount: 5,
			items:       []int{8},
			expected:    []int{1, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SpliceArr(tt.arr, tt.start, tt.deleteCount, tt.items...)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("SpliceArr() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTake(t *testing.T) {
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
			name: "no take",
			arr:  []int{1, 2, 3},
			n:    0,
			want: nil,
		},
		{
			name: "take one",
			arr:  []int{1, 2, 3},
			n:    1,
			want: []int{1},
		},
		{
			name: "take two",
			arr:  []int{1, 2, 3},
			n:    2,
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Take(tt.arr, tt.n)

			if len(got) != len(tt.want) {
				t.Errorf("Take() = %v, want %v", got, tt.want)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Take() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestReverseArr(t *testing.T) {
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
			name: "single element",
			arr:  []int{1},
			want: []int{1},
		},
		{
			name: "multiple elements",
			arr:  []int{1, 2, 3},
			want: []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseArr(tt.arr)

			if len(got) != len(tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Reverse() = %v, want %v", got, tt.want)
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

func TestSort(t *testing.T) {
	intArray := []int{5, 2, 9, 1, 5, 6}
	Sort(intArray)

	expected := []int{1, 2, 5, 5, 6, 9}
	for i, v := range intArray {
		if v != expected[i] {
			t.Errorf("expected %v, got %v", expected, intArray)
		}
	}

	strArray := []string{"apple", "orange", "banana", "grape"}
	Sort(strArray)

	expected2 := []string{"apple", "banana", "grape", "orange"}
	for i, v := range strArray {
		if v != expected2[i] {
			t.Errorf("expected %v, got %v", expected2, strArray)
		}
	}

	floatArray := []float64{5.2, 3.1, 7.4, 1.9}
	Sort(floatArray)

	expected3 := []float64{1.9, 3.1, 5.2, 7.4}
	for i, v := range floatArray {
		if v != expected3[i] {
			t.Errorf("expected %v, got %v", expected3, floatArray)
		}
	}

	emptyArray := []int{}
	Sort(emptyArray)

	if len(emptyArray) != 0 {
		t.Errorf("expected empty array, got %v", emptyArray)
	}

	singleElementArray := []int{42}
	Sort(singleElementArray)

	expected4 := []int{42}
	for i, v := range singleElementArray {
		if v != expected4[i] {
			t.Errorf("expected %v, got %v", expected4, singleElementArray)
		}
	}
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		input    any
		expected []any
	}{
		{
			input: []any{
				[]any{1, 2, []any{3, 4}, 5},
				6, []any{7, []any{8, 9}},
			},
			expected: []any{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			input:    []any{1, 2, 3},
			expected: []any{1, 2, 3},
		},
		{
			input:    []any{[]any{1, 2}, 3, []any{4, 5}},
			expected: []any{1, 2, 3, 4, 5},
		},
		{
			input:    []any{[]any{}, []any{1, 2}},
			expected: []any{1, 2},
		},
		{
			input:    []any{1, []any{2, []any{3, 4}}, 5},
			expected: []any{1, 2, 3, 4, 5},
		},
		{
			input:    []any{},
			expected: []any{},
		},
	}

	for _, test := range tests {
		t.Run("Flatten", func(t *testing.T) {
			result := Flatten(test.input)
			if len(result) == 0 && len(test.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func _isDifferent[T comparable](arr1, arr2 []T) bool {
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return true
		}
	}
	return false
}

func TestShuffle(t *testing.T) {
	intArr_1 := []int{1, 2, 3, 4, 5}
	intArr_2 := []int{1, 2, 3, 4, 5}
	Shuffle(intArr_2)

	if len(intArr_1) != len(intArr_2) {
		t.Errorf("Expected shuffled array to have same length as original, but got length %d", len(intArr_2))
	}

	if !_isDifferent(intArr_1, intArr_2) {
		t.Errorf("Expected shuffled array to be different from original")
	}

	emptyArr := []int{}

	if len(emptyArr) != 0 {
		t.Errorf("Expected shuffled empty array to have length 0, but got length %d", len(emptyArr))
	}

	strArr_1 := []string{"apple", "banana", "cherry", "date"}
	strArr_2 := []string{"apple", "banana", "cherry", "date"}
	Shuffle(strArr_2)

	if len(strArr_1) != len(strArr_2) {
		t.Errorf("Expected shuffled array to have same length as original, but got length %d", len(strArr_2))
	}
	if !_isDifferent(strArr_1, strArr_2) {
		t.Errorf("Expected shuffled array to be different from original")
	}
}

func TestIsEqualArr(t *testing.T) {
	intArr1 := []int{1, 2, 3, 4}
	intArr2 := []int{1, 2, 3, 4}
	if !IsEqualArr(intArr1, intArr2) {
		t.Errorf("Expected %v and %v to be equal", intArr1, intArr2)
	}

	intArr3 := []int{4, 3, 2, 1}
	if IsEqualArr(intArr1, intArr3) {
		t.Errorf("Expected %v and %v to be different", intArr1, intArr3)
	}

	strArr1 := []string{"apple", "banana", "cherry"}
	strArr2 := []string{"apple", "banana", "cherry"}
	if !IsEqualArr(strArr1, strArr2) {
		t.Errorf("Expected %v and %v to be equal", strArr1, strArr2)
	}

	strArr3 := []string{"cherry", "banana", "apple"}
	if IsEqualArr(strArr1, strArr3) {
		t.Errorf("Expected %v and %v to be different", strArr1, strArr3)
	}

	emptyArr1 := []int{}
	emptyArr2 := []int{}
	if !IsEqualArr(emptyArr1, emptyArr2) {
		t.Errorf("Expected %v and %v to be equal", emptyArr1, emptyArr2)
	}

	unequalArr := []int{1, 2}
	if IsEqualArr(intArr1, unequalArr) {
		t.Errorf("Expected %v and %v to be different", intArr1, unequalArr)
	}
}