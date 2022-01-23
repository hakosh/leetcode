package rotate_image

import (
	"reflect"
	"testing"
)

type Test struct {
	in  [][]int
	out [][]int
}

var square1 = [][]int{
	{1, 2, 3, 4, 5},
	{6, 7, 8, 9, 10},
	{11, 12, 13, 14, 15},
	{16, 17, 18, 19, 20},
	{21, 22, 23, 24, 25},
}

var square1R = [][]int{
	{21, 16, 11, 6, 1},
	{22, 17, 12, 7, 2},
	{23, 18, 13, 8, 3},
	{24, 19, 14, 9, 4},
	{25, 20, 15, 10, 5},
}

var tests = []Test{
	{[][]int{{1}}, [][]int{{1}}},
	{square1, square1R},
}

func TestRotate(t *testing.T) {
	for i, test := range tests {
		rotate(test.in)

		if !reflect.DeepEqual(test.in, test.out) {
			t.Errorf("For test %d wanted %v, got %v", i, test.out, test.in)
		}
	}
}

func BenchmarkRotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotate(square1)
	}
}
