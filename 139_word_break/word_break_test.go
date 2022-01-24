package word_break

import "testing"

type Test struct {
	s    string
	dict []string
	out  bool
}

var tests = []Test{
	{"leetcode", []string{"leet", "code"}, true},
	{"cat", []string{"c", "airplane", "book"}, false},
	{"abcdef", []string{"abcde", "ef", "abc", "a", "d"}, true},
	{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
}

func TestWordBreak(t *testing.T) {
	for _, test := range tests {
		if res := wordBreak(test.s, test.dict); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.s, test.dict, test.out, res)
		}
	}
}
