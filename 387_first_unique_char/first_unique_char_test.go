package first_unique_char

import (
	"testing"
)

type Test struct {
	s    string
	want int
}

var cases = []Test{
	//{s: "leetcode", want: 0},
	//{s: "loveleetcode", want: 2},
	{s: "aabb", want: -1},
	//{s: "acaadcad", want: -1},
	//{s: "dddccdbba", want: 8},
}

func TestFirstUniqueChar(t *testing.T) {
	for _, test := range cases {
		if res := firstUniqChar(test.s); res != test.want {
			t.Errorf("For %s wanted %v, got %v\n", test.s, test.want, res)
		}
	}
}
