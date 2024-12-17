package gouse

import "testing"

/* Testing min max */

func TestMin(t *testing.T) {
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

/* Testing intersection */

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

/* Testing callback functions */

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

/* Testing utilities */

func TestExist(t *testing.T) {
	if !IncludesArr[int]([]int{1, 2, 3}, 1) {
		t.Errorf("Exist[int]([]int{1, 2, 3}, 1) = false, want true")
	}
}

func TestEqual(t *testing.T) {
	if !IncludesArr[int]([]int{1, 2, 3}, 1) {
		t.Errorf("Equal[int](1, 1) = false, want true")
	}
}
