package time_needed_to_inform_all_employees

import "testing"

type Test struct {
	n          int
	headID     int
	manager    []int
	informTime []int
	out        int
}

var tests = []Test{
	{6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 4, 0, 0, 0}, 4},
	{6, 2, []int{1, 2, -1, 2, 1, 2}, []int{0, 2, 1, 0, 0, 0}, 3},
	{1, 0, []int{-1}, []int{9}, 0},
	{6, 1, []int{1, -1}, []int{0, 0}, 0},
	{6, 1, []int{1, -1}, []int{0, 100}, 100},
	{6, 0, []int{-1, 0, 1, 2, 3}, []int{1, 1, 1, 1, 0}, 4},
	{11, 4, []int{5, 9, 6, 10, -1, 8, 9, 1, 9, 3, 4}, []int{0, 213, 0, 253, 686, 170, 975, 0, 261, 309, 337}, 2560},
}

func TestNumOfMinutes(t *testing.T) {
	for _, test := range tests {
		if res := numOfMinutes(test.n, test.headID, test.manager, test.informTime); res != test.out {
			t.Errorf("For %v expected %v, got %v", test, test.out, res)
		}
	}
}
