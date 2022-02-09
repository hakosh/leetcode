package course_schedule

import "testing"

type Test struct {
	num           int
	prerequisites [][]int
	out           bool
}

var tests = []Test{
	{0, [][]int{{}}, true},
	{1, [][]int{}, true},
	{2, [][]int{{0, 1}}, true},
	{2, [][]int{{0, 1}, {1, 0}}, false},
	{6, [][]int{{1, 0}, {2, 1}, {2, 5}, {0, 3}, {4, 3}, {3, 5}, {4, 5}}, true},
	{7, [][]int{{0, 3}, {1, 0}, {2, 1}, {4, 5}, {6, 4}, {5, 6}}, false},
	{8, [][]int{{0, 3}, {1, 0}, {2, 1}, {4, 5}, {6, 4}, {5, 6}, {5, 7}}, false},
	{3, [][]int{{1, 0}, {1, 2}, {0, 1}}, false},
}

func TestCanFinish(t *testing.T) {
	for _, test := range tests {
		if res := canFinish(test.num, test.prerequisites); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.num, test.prerequisites, test.out, res)
		}
	}
}
