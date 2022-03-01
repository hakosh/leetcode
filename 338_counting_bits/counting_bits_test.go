package counting_bits

import (
	"reflect"
	"testing"
)

type Test struct {
	n   int
	out []int
}

var tests = []Test{
	{0, []int{0}},
	{2, []int{0, 1, 1}},
	{5, []int{0, 1, 1, 2, 1, 2}},
}

func TestCountBits(t *testing.T) {
	for _, test := range tests {
		if res := countBits(test.n); !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v expected %v, got %v", test.n, test.out, res)
		}
	}
}
