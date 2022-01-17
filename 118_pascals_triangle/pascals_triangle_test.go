package pascals_triangle

import (
	"reflect"
	"testing"
)

type Test struct {
	n   int
	out [][]int
}

var cases = []Test{
	{n: 1, out: [][]int{{1}}},
	{n: 2, out: [][]int{{1}, {1, 1}}},
	{n: 3, out: [][]int{{1}, {1, 1}, {1, 2, 1}}},
	{n: 4, out: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}},
	{n: 5, out: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
}

func TestGenerate(t *testing.T) {
	for _, test := range cases {
		res := generate(test.n)

		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("For %v wanted %v, got %v\n", test.n, test.out, res)
		}
	}
}
