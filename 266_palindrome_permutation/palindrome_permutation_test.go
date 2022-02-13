package palindrome_permutation

import "testing"

type Test struct {
	in  string
	out bool
}

var tests = []Test{
	{"carerac", true},
	{"aab", true},
	{"code", false},
	{"a", true},
	{"aabbccx", true},
	{"ayabbccx", false},
}

func TestCanPermutatePalindrome(t *testing.T) {
	for _, test := range tests {
		if res := canPermutePalindrome(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
