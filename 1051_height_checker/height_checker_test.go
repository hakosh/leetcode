package height_checker

import (
	"testing"
)

type Test struct {
	in   []int
	want int
}

var cases = []Test{
	{in: []int{1, 1, 4, 2, 1, 3}, want: 3},
	{in: []int{5, 1, 2, 3, 4}, want: 5},
	{in: []int{1, 2, 3, 4, 5}, want: 0},
}

func TestHeightChecker(t *testing.T) {
	for _, test := range cases {
		if res := heightChecker(test.in); res != test.want {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.want, res)
		}
	}
}
