package reshape_the_matrix

import (
	"reflect"
	"testing"
)

type Test struct {
	mat [][]int
	r   int
	c   int
	out [][]int
}

var cases = []Test{
	{mat: [][]int{{1, 2}, {3, 4}}, r: 1, c: 4, out: [][]int{{1, 2, 3, 4}}},
	{mat: [][]int{{1, 2}, {3, 4}}, r: 4, c: 1, out: [][]int{{1}, {2}, {3}, {4}}},
	{mat: [][]int{{1, 2}, {3, 4}}, r: 2, c: 4, out: [][]int{{1, 2}, {3, 4}}},
	{mat: [][]int{{1, 2, 3, 4}}, r: 2, c: 2, out: [][]int{{1, 2}, {3, 4}}},
}

func TestIntersect(t *testing.T) {
	for _, test := range cases {
		res := matrixReshape(test.mat, test.r, test.c)

		if !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v and %v and %v wanted %v, got %v\n", test.mat, test.r, test.c, test.out, res)
		}
	}
}
