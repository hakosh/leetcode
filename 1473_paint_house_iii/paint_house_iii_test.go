package paint_house_iii

import "testing"

type Test struct {
	houses []int
	cost   [][]int
	m      int
	n      int
	target int
	out    int
}

var tests = []Test{
	{[]int{0, 0, 0, 0, 0}, [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}, 5, 2, 3, 9},
	{[]int{0, 2, 1, 2, 0}, [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}, 5, 2, 3, 11},
	{[]int{3, 1, 2, 3}, [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 4, 3, 3, -1},
	{[]int{0, 0, 0, 1}, [][]int{{1, 5}, {4, 1}, {1, 3}, {4, 4}}, 4, 2, 4, 12},
}

func TestMinCost(t *testing.T) {
	for i, test := range tests {
		if res := minCost(test.houses, test.cost, test.m, test.n, test.target); res != test.out {
			t.Errorf("For case %d expected %d, got %d", i, test.out, res)
		}
	}
}
