package valid_palindrome_ii

import "testing"

type Test struct {
	in  string
	out bool
}

var tests = []Test{
	{"raceacar", true},
	{" ", true},
	{"xd", true},
	{"eeccccbebaeeabebccceea", false},
	{"aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga", true},
	{"axbbxax", true},
	{"axxax", true},
	{"acxcybycxcxa", true},
	{"abccbxa", true},
	{"majom", false},
	{"abca", true},
	{"aba", true},
	{"11XX11", true},
	{"12XX11", false},
}

func TestValidPalindrome(t *testing.T) {
	for _, test := range tests {
		if res := validPalindrome(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
