package champagne_tower

import "testing"

type Test struct {
	poured     int
	queryRow   int
	queryGlass int
	out        float64
}

var tests = []Test{
	{0, 0, 0, 0},
	{2, 0, 0, 1},
	{30, 0, 0, 1},
	{0, 99, 98, 0},
	{1, 0, 0, 1.0},
	{2, 1, 1, 0.5},
	{4, 2, 0, 0.25},
	{20, 5, 1, 0.75},
	{100000009, 33, 17, 1.0},
}

func TestChampagneTower(t *testing.T) {
	for _, test := range tests {
		if res := champagneTower(test.poured, test.queryRow, test.queryGlass); res != test.out {
			t.Errorf("For %v, %v, and %v expected %v, got %v", test.poured, test.queryRow, test.queryGlass, test.out, res)
		}
	}
}
