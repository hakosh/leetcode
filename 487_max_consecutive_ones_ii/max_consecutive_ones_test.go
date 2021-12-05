package max_consecutive_ones_ii

import (
	"testing"
)

type Test struct {
	in   []int
	want int
}

var cases = []Test{
	{in: []int{1, 1, 0, 1}, want: 4},
	{in: []int{1, 1, 1, 1}, want: 4},
	{in: []int{1, 0, 1, 1, 0}, want: 4},
	{in: []int{1, 0, 1, 1, 0, 1}, want: 4},
	{in: []int{1, 0, 0, 1, 1, 0, 0, 1}, want: 3},
	{in: []int{1}, want: 1},
	{in: []int{0}, want: 1},
	{in: []int{0, 0}, want: 1},
	{in: []int{1, 1, 1, 0}, want: 4},
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	for _, test := range cases {
		if res := findMaxConsecutiveOnes(test.in); res != test.want {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.want, res)
		}
	}
}
