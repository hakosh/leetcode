package ransom_note

import (
	"testing"
)

type Test struct {
	note     string
	magazine string
	want     bool
}

var cases = []Test{
	{note: "a", magazine: "b", want: false},
	{note: "aa", magazine: "ab", want: false},
	{note: "aa", magazine: "aab", want: true},
	{note: "liba", magazine: "banana textile", want: true},
}

func TestCanConstruct(t *testing.T) {
	for _, test := range cases {
		if res := canConstruct(test.note, test.magazine); res != test.want {
			t.Errorf("For %s and %s wanted %v, got %v\n", test.note, test.magazine, test.want, res)
		}
	}
}
