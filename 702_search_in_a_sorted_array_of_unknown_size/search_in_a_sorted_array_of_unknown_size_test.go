package search_in_a_sorted_array_of_unknown_size

import "testing"

type Test struct {
	nums   []int
	target int
	out    int
}

var tests = []Test{
	{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
	{[]int{3}, 3, 0},
	{[]int{3}, 2, -1},
}

func TestSearch(t *testing.T) {
	for _, test := range tests {
		reader := ArrayReader{items: test.nums}

		if res := search(reader, test.target); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.nums, test.out, res)
		}
	}
}
