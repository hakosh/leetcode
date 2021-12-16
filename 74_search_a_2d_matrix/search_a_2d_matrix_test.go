package search_a_2d_matrix

import (
	"testing"
)

type Test struct {
	mat    [][]int
	target int
	out    bool
}

var cases = []Test{
	{mat: [][]int{{1}}, target: 1, out: true},
	{mat: [][]int{{1, 2}}, target: 2, out: true},
	{mat: [][]int{{1, 2}, {3, 4}}, target: 2, out: true},
	{mat: [][]int{{1, 2}, {3, 4}}, target: 5, out: false},
	{mat: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, target: 3, out: true},
	{mat: [][]int{{1, 3, 5, 7, 10}, {11, 16, 20, 23, 30}}, target: 16, out: true},
}

func TestIntersect(t *testing.T) {
	for _, test := range cases {
		res := searchMatrix(test.mat, test.target)

		if res != test.out {
			t.Errorf("For %v and %v wanted %v, got %v\n", test.mat, test.target, test.out, res)
		}
	}
}
