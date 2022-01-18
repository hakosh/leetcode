package largest_perimeter_triangle

import (
	"testing"
)

type Test struct {
	in  []int
	out int
}

var tests = []Test{
	{[]int{2, 1, 2}, 5},
	{[]int{1, 2, 1}, 0},
	{[]int{2, 1, 2, 10, 8, 4, 1, 5, 8, 7}, 26},
}

func TestDecodeString(t *testing.T) {
	for _, test := range tests {
		if res := largestPerimeter(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
