package flood_fill

import (
	"reflect"
	"testing"
)

type Test struct {
	image    [][]int
	sr       int
	sc       int
	newColor int
	out      [][]int
}

var image1 = [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
var out1 = [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}

var image2 = [][]int{{0, 0, 0}, {0, 1, 1}}
var out2 = [][]int{{0, 0, 0}, {0, 1, 1}}

var tests = []Test{
	{image1, 1, 1, 2, out1},
	{image2, 1, 1, 1, out2},
}

func TestFloodFill(t *testing.T) {
	for i, test := range tests {
		if out := floodFill(test.image, test.sr, test.sc, test.newColor); !reflect.DeepEqual(out, test.out) {
			t.Errorf("For case #%d expected %v, got %v", i, test.out, out)
		}
	}
}
