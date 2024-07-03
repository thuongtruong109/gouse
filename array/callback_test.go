package array

import (
	"testing"
)

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

