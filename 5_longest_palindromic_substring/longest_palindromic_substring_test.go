package longest_palindromic_substring

import "testing"

type Test struct {
	s   string
	out string
}

var tests = []Test{
	{"babad", "bab"},
	{"xyracecar", "racecar"},
	{"lkfhjsfldksuifdusykyyyydklmsracecardlsfkadlksj", "racecar"},
}

func TestLongestPalindrome(t *testing.T) {
	for _, test := range tests {
		if res := longestPalindrome(test.s); res != test.out {
			t.Errorf("For %q expected %q, got %q", test.s, test.out, res)
		}
	}
}

func BenchmarkLongestPalindrome(b *testing.B) {
	s := "lkfhjsfldksuifdusykyyyydklmsracecardlsfkadlksj"

	for i := 0; i < b.N; i++ {
		longestPalindrome(s)
	}
}
