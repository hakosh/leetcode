package search_in_rotated_sorted_array

import "testing"

type Test struct {
	nums   []int
	target int
	out    int
}

var tests = []Test{
	{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
	{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
	{[]int{1}, 1, 0},
	{[]int{0, 1, 2, 3, 4, 5, 6}, 4, 4},
	{[]int{1, 2, 3, 4, 5, 6, 0}, 4, 3},
}

func TestSearch(t *testing.T) {
	for _, test := range tests {
		if res := search(test.nums, test.target); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.nums, test.target, test.out, res)
		}
	}
}
