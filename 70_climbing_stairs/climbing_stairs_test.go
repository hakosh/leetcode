package climbing_stairs

import (
	"testing"
)

type Test struct {
	n   int
	out int
}

var cases = []Test{
	{n: 2, out: 2},
	{n: 3, out: 3},
	{n: 45, out: 1836311903},
}

func TestIntersect(t *testing.T) {
	for _, test := range cases {
		res := climbStairs(test.n)

		if res != test.out {
			t.Errorf("For %v wanted %v, got %v\n", test.n, test.out, res)
		}
	}
}
