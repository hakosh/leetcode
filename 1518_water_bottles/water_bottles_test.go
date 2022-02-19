package water_bottles

import "testing"

type Test struct {
	numBottles  int
	numExchange int
	out         int
}

var tests = []Test{
	{15, 4, 19},
}

func TestNumWaterBottles(t *testing.T) {
	for _, test := range tests {
		if res := numWaterBottles(test.numBottles, test.numExchange); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.numBottles, test.numExchange, test.out, res)
		}
	}
}
