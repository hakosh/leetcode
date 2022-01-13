package matrix

import (
	"reflect"
	"testing"
)

type Test struct {
	in  [][]int
	out [][]int
}

var tests = []Test{
	{
		[][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		[][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
	},
	{
		[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
	},
	{
		[][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1, 0}},
		[][]int{{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
	},
	{
		[][]int{{0, 1, 0, 1, 1}, {1, 1, 0, 0, 1}, {0, 0, 0, 1, 0}, {1, 0, 1, 1, 1}, {1, 0, 0, 0, 1}},
		[][]int{{0, 1, 0, 1, 2}, {1, 1, 0, 0, 1}, {0, 0, 0, 1, 0}, {1, 0, 1, 1, 1}, {1, 0, 0, 0, 1}},
	},
}

func TestUpdateMatrix(t *testing.T) {
	for _, test := range tests {
		if res := updateMatrix(test.in); !reflect.DeepEqual(test.out, res) {
			t.Errorf("For %v wanted %v, got %v", test.in, test.out, res)
		}
	}
}

func BenchmarkUpdateMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := [][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1, 0}}
		updateMatrix(in)
	}
}
