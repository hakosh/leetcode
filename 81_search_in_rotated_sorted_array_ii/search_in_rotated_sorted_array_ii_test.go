package search_in_rotated_sorted_array_ii

import "testing"

type Test struct {
	nums   []int
	target int
	want   bool
}

var tests = []Test{
	{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
}

func TestSearch(t *testing.T) {
	for _, test := range tests {
		if res := search(test.nums, test.target); res != test.want {
			t.Errorf("for %v and %v wanted %v, got %v", test.nums, test.target, test.want, res)
		}
	}
}
