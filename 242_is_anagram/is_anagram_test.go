package is_anagram

import (
	"testing"
)

type Test struct {
	s    string
	t    string
	want bool
}

var cases = []Test{
	{s: "anagram", t: "nagaram", want: true},
	{s: "aaiii", t: "aaiib", want: false},
	{s: "rat", t: "car", want: false},
	{s: "a", t: "b", want: false},
	{s: "ab", t: "a", want: false},
}

func TestIsAnagram(t *testing.T) {
	for _, test := range cases {
		if res := isAnagram(test.s, test.t); res != test.want {
			t.Errorf("For %s and %s wanted %v, got %v\n", test.t, test.s, test.want, res)
		}
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	t := "anagram"
	s := "nagaram"

	for i := 0; i < b.N; i++ {
		isAnagram(s, t)
	}
}
