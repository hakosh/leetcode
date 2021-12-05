package third_maximum_number

import (
	"math"
	"testing"
)

type Test struct {
	in   []int
	want int
}

var cases = []Test{
	{in: []int{3, 2, 1}, want: 1},
	{in: []int{3, 3, 2, 1}, want: 1},
	{in: []int{3, 3, 2, 1, 3, 1}, want: 1},
	{in: []int{1, 2, 2, 2, 2}, want: 2},
	{in: []int{1, 1, 2, 2, 3, 3}, want: 1},
	{in: []int{1, 2}, want: 2},
	{in: []int{-5, -5, -5, 2}, want: 2},
	{in: []int{-5, -5, -5, 2, -10}, want: -10},
	{in: []int{2, 2, 3, 1}, want: 1},
	{in: []int{2, 2, 3, 1, -1, 0}, want: 1},
	{in: []int{1}, want: 1},
	{in: []int{1, 2, math.MinInt32}, want: math.MinInt32},
}

func TestThirdMax(t *testing.T) {
	for _, test := range cases {
		if res := thirdMax(test.in); res != test.want {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.want, res)
		}
	}
}
