package rotting_oranges

import "testing"

type Test struct {
	in  [][]int
	out int
}

var grid1 = [][]int{
	{2, 1, 1},
	{1, 1, 0},
	{0, 1, 1},
}

var grid2 = [][]int{
	{2, 1, 1},
	{0, 1, 1},
	{1, 0, 1},
}

var grid3 = [][]int{
	{2, 1, 1},
	{0, 1, 1},
	{1, 2, 1},
}

var grid4 = [][]int{
	{0, 2},
}

var tests = []Test{
	{grid1, 4},
	{grid2, -1},
	{grid3, 2},
	{grid4, 0},
	{[][]int{{1}}, -1},
	{[][]int{{0}}, 0},
}

func TestOrangesRotting(t *testing.T) {
	for _, test := range tests {
		if res := orangesRotting(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
