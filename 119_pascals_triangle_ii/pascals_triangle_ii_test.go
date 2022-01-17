package pascals_triangle_ii

import (
	"reflect"
	"testing"
)

type Test struct {
	n   int
	out []int
}

var cases = []Test{
	{n: 0, out: []int{1}},
	{n: 1, out: []int{1, 1}},
	{n: 2, out: []int{1, 2, 1}},
	{n: 3, out: []int{1, 3, 3, 1}},
	{n: 4, out: []int{1, 4, 6, 4, 1}},
}

func TestGetRow(t *testing.T) {
	for _, test := range cases {
		res := getRow(test.n)

		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("For %v wanted %v, got %v\n", test.n, test.out, res)
		}
	}
}

func BenchmarkGetRow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getRow(50)
	}
}
