package kth_symbol_in_grammar

import (
	"testing"
)

type Test struct {
	n   int
	k   int
	out int
}

var tests = []Test{
	{1, 1, 0},
	{2, 1, 0},
	{2, 2, 1},
	{3, 1, 0},
	{5, 6, 0},
	{30, 434991989, 0},
}

func TestMergeTwoLists(t *testing.T) {
	for _, test := range tests {
		if res := kthGrammar(test.n, test.k); res != test.out {
			t.Errorf("For %v %v wanted %v, got %v\n", test.n, test.k, test.out, res)
		}
	}
}
