package contains_duplicate

import (
	"testing"
)

type Test struct {
	in []int
	want bool
}

var cases = []Test{
	{ in: []int{}, want: false },
	{ in: []int{1, 2, 3, 4}, want: false},
	{ in: []int{0, 1, 0, 3, 12}, want: true },
	{ in: []int{1,1,1,3,3,4,3,2,4,2}, want: true},
	{ in: []int{-24500, -24500}, want: true},
}

func TestMoveZeroes(t *testing.T) {
	for _, test := range cases {
		result := containsDuplicate(test.in)

		if result != test.want {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.want, result)
		}
	}
}
