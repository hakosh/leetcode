package find_pivot_index

import "testing"

type Test struct {
	in  []int
	out int
}

var tests = []Test{
	{[]int{1, 7, 3, 6, 5, 6}, 3},
	{[]int{2, 1, -1}, 0},
	{[]int{1}, 0},
	{[]int{0}, 0},
	{[]int{-1, -1, -1, -1, -1, 1}, -1},
	{[]int{-1, -1, 0, 0, -1, -1}, 2},
	{[]int{-1, -1, 0, 1, -1, -1}, 1},
}

func TestPivotIndex(t *testing.T) {
	for _, test := range tests {
		if res := pivotIndex(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
