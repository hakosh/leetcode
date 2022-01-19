package min_cost_of_climbing_stairs

import (
	"testing"
)

type Test struct {
	in  []int
	out int
}

var cases = []Test{
	{[]int{10, 15, 20}, 15},
	{[]int{5, 15}, 5},
	{[]int{0, 0, 0, 1}, 0},
	{[]int{0, 0, 1, 1}, 1},
	{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
}

func TestMinCostClimbingStairs(t *testing.T) {
	for _, test := range cases {
		res := minCostClimbingStairs(test.in)

		if res != test.out {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.out, res)
		}
	}
}
