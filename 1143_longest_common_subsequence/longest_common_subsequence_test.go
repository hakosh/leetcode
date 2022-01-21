package longest_common_subsequence

import "testing"

type Test struct {
	text1 string
	text2 string
	out   int
}

var tests = []Test{
	{"abcde", "ace", 3},
	{"abc", "abc", 3},
	{"abc", "def", 0},
	{"abcde", "ecd", 2},
}

func TestLongestCommonSubsequence(t *testing.T) {
	for _, test := range tests {
		res := longestCommonSubsequence(test.text1, test.text2)

		if res != test.out {
			t.Errorf("For %v and %v wanted %v, got %v\n", test.text1, test.text2, test.out, res)
		}
	}
}

func BenchmarkLongestCommonSubsequence(b *testing.B) {
	t1 := "goinactionlongwordballpointpenpenguin"
	t2 := "tpenpeng"

	for i := 0; i < b.N; i++ {
		longestCommonSubsequence(t1, t2)
	}
}
