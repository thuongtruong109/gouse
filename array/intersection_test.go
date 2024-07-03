package array

import (
	"testing"
)

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