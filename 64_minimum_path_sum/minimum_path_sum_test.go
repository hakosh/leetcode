package minimum_path_sum

import "testing"

type Test struct {
	grid [][]int
	out  int
}

var grid1 = [][]int{
	{1, 3, 1},
	{1, 5, 1},
	{4, 2, 1},
}

var grid2 = [][]int{
	{1, 2, 3},
	{4, 5, 6},
}

var tests = []Test{
	{grid1, 7},
	{grid2, 12},
}

func TestUniquePathsWithObstacles(t *testing.T) {
	for _, test := range tests {
		if res := minPathSum(test.grid); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.grid, test.out, res)
		}
	}
}
