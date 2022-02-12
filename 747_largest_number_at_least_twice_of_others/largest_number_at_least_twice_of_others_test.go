package largest_number_at_least_twice_of_others

import "testing"

type Test struct {
	in  []int
	out int
}

var tests = []Test{
	{[]int{3, 6, 1, 0}, 1},
	{[]int{1, 2, 3, 4}, -1},
	{[]int{1}, 0},
	{[]int{30, 60}, 1},
	{[]int{30, 59}, -1},
	{[]int{29, 59}, 1},
	{[]int{60, 30}, 0},
	{[]int{59, 30}, -1},
	{[]int{2, 3, 4}, -1},
}

func TestDominantIndex(t *testing.T) {
	for _, test := range tests {
		if res := dominantIndex(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
