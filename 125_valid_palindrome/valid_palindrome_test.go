package valid_palindrome

import "testing"

type Test struct {
	in  string
	out bool
}

var tests = []Test{
	{"A man, a plan, a canal: Panama", true},
	{"race a car", false},
	{" ", true},
	{"majom", false},
	{"11XX11", true},
	{"12XX11", false},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range tests {
		if res := isPalindrome(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
