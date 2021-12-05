package find_all_numbers_disappeared_in_an_array

import (
	"reflect"
	"testing"
)

type Test struct {
	in   []int
	want []int
}

var cases = []Test{
	{in: []int{4, 3, 2, 7, 8, 2, 3, 1}, want: []int{5, 6}},
	{in: []int{1, 1}, want: []int{2}},
}

func TestFindDisappearedNumbers(t *testing.T) {
	for _, test := range cases {
		res := findDisappearedNumbers(test.in)

		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.want, res)
		}
	}
}
