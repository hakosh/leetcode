package sort_integers_by_the_number_of_1_bits

import (
	"reflect"
	"testing"
)

type Test struct {
	in  []int
	out []int
}

var tests = []Test{
	{[]int{1}, []int{1}},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, []int{0, 1, 2, 4, 8, 3, 5, 6, 7}},
	{[]int{1024, 512, 256, 128, 32, 64, 16, 8, 4, 2, 1}, []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}},
}

func TestSortByBits(t *testing.T) {
	for _, test := range tests {
		in := make([]int, len(test.in))
		copy(in, test.in)

		if res := sortByBits(test.in); !reflect.DeepEqual(test.out, res) {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
