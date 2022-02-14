package binary_search

import "testing"

type Test struct {
	nums   []int
	target int
	out    int
}

var tests = []Test{
	{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
	{[]int{-1, 0, 3, 5, 9, 12, 14}, 9, 4},
	{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	{[]int{1}, 1, 0},
	{[]int{1, 2}, 1, 0},
	{[]int{1, 2, 3}, 1, 0},
	{[]int{1, 2, 3}, 3, 2},
}

func TestSearch(t *testing.T) {
	for _, test := range tests {
		if res := search(test.nums, test.target); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.nums, test.target, test.out, res)
		}
	}
}
