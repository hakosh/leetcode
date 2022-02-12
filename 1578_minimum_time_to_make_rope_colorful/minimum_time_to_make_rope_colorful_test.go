package minimum_time_to_make_rope_colorful

import "testing"

type Test struct {
	colors     string
	neededTime []int
	out        int
}

var tests = []Test{
	{"abaac", []int{1, 2, 3, 4, 5}, 3},
	{"abaaac", []int{1, 2, 3, 4, 5, 6}, 7},
	{"abc", []int{1, 2, 3}, 0},
	{"a", []int{1}, 0},
	{"bbbbbb", []int{2, 2, 2, 2, 2, 2}, 10},
	{"bbxbbbx", []int{2, 3, 2, 3, 2, 4, 3}, 7},
}

func TestMinCost(t *testing.T) {
	for _, test := range tests {
		if res := minCost(test.colors, test.neededTime); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.colors, test.neededTime, test.out, res)
		}
	}
}
