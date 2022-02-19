package find_three_consecutive_integers_that_sum_to_a_given_number

import (
	"reflect"
	"testing"
)

type Test struct {
	num int64
	out []int64
}

var tests = []Test{
	{4, []int64{}},
	{0, []int64{-1, 0, 1}},
	{33, []int64{10, 11, 12}},
}

func TestSumOfThree(t *testing.T) {
	for _, test := range tests {
		if res := sumOfThree(test.num); !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v expected %v, got %v", test.num, test.out, res)
		}
	}
}
