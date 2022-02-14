package find_n_unique_integers_sum_up_to_zero

import (
	"reflect"
	"testing"
)

type Test struct {
	n   int
	out []int
}

var tests = []Test{
	{5, []int{1, -1, 2, -2, 0}},
	{0, []int{}},
	{1, []int{0}},
}

func TestSumZero(t *testing.T) {
	for _, test := range tests {
		if res := sumZero(test.n); !reflect.DeepEqual(test.out, res) {
			t.Errorf("For %v expected %v, got %v", test.n, test.out, res)
		}
	}
}
