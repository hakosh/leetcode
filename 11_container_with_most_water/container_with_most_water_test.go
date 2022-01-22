package container_with_most_water

import "testing"

type Test struct {
	in  []int
	out int
}

var tests = []Test{
	{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	{[]int{1, 3, 2, 5, 25, 24, 5}, 24},
	{[]int{1, 1}, 1},
	{[]int{1, 2, 1}, 2},
}

func TestMaxArea(t *testing.T) {
	for _, test := range tests {
		if res := maxArea(test.in); res != test.out {
			t.Errorf("For %v wanted %v, got %v", test.in, test.out, res)
		}
	}
}
