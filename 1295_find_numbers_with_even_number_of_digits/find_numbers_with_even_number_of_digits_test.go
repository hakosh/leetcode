package find_numbers_with_even_number_of_digits

import "testing"

type Test struct {
	in  []int
	out int
}

var tests = []Test{
	{in: []int{12, 345, 2, 6, 7896}, out: 2},
	{in: []int{555, 901, 482, 1771}, out: 1},
	{in: []int{}, out: 0},
	{in: []int{100000}, out: 1},
}

func TestFindNumbers(t *testing.T) {
	for _, test := range tests {
		res := findNumbers(test.in)

		if res != test.out {
			t.Errorf("For %v wanted %d, got %d\n", test.in, test.out, res)
		}
	}
}
