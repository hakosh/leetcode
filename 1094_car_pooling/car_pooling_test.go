package car_pooling

import "testing"

type Test struct {
	trips    [][]int
	capacity int
	out      bool
}

var tests = []Test{
	{[][]int{{2, 1, 5}, {3, 3, 7}}, 4, false},
	{[][]int{{2, 1, 5}, {3, 3, 7}}, 5, true},
}

func TestCarPooling(t *testing.T) {
	for _, test := range tests {
		if res := carPooling(test.trips, test.capacity); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.trips, test.capacity, test.out, res)
		}
	}
}
