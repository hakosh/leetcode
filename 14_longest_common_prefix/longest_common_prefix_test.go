package longest_common_prefix

import (
	"fmt"
	"testing"
)

type Test struct {
	in  []string
	out string
}

var cases = []Test{
	{in: []string{"flow", "flower", "flights"}, out: "fl"},
	{in: []string{"dog", "racecar", "car"}, out: ""},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, test := range cases {
		got := longestCommonPrefix(test.in)

		if got != test.out {
			t.Errorf("For %v, expected %s, got %s\n", test.in, test.out, got)
		}
	}
}

func Example_longestCommonPrefix() {
	fmt.Println(longestCommonPrefix([]string{"dog", "dope"}))
	// Output: do
}
