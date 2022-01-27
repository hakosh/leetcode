package unique_paths_ii

import "testing"

type Test struct {
	grid [][]int
	out  int
}

var grid1 = [][]int{
	{0, 0, 0},
	{0, 1, 0},
	{0, 0, 0},
}

var grid2 = [][]int{
	{0, 1},
	{0, 0},
}

var tests = []Test{
	{grid1, 2},
	{grid2, 1},
}

func TestUniquePathsWithObstacles(t *testing.T) {
	for _, test := range tests {
		if res := uniquePathsWithObstacles(test.grid); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.grid, test.out, res)
		}
	}
}
