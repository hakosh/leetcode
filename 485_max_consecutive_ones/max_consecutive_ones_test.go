package max_consecutive_ones

import "testing"

type Test struct {
	in  []int
	out int
}

var tests = []Test{
	{in: []int{1, 1, 0, 1, 1, 1}, out: 3},
	{in: []int{1, 0, 1, 1, 0, 1}, out: 2},
	{in: []int{}, out: 0},
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	for _, test := range tests {
		res := findMaxConsecutiveOnes(test.in)

		if res != test.out {
			t.Errorf("For %v wanted %d, got %d\n", test.in, test.out, res)
		}
	}
}