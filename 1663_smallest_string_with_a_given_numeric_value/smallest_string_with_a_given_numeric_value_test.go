package smallest_string_with_a_given_numeric_value

import "testing"

type Test struct {
	n    int
	k    int
	want string
}

var tests = []Test{
	{3, 27, "aay"},
	{5, 73, "aaszz"},
}

func TestGetSmallestString(t *testing.T) {
	for _, test := range tests {
		if res := getSmallestString(test.n, test.k); res != test.want {
			t.Errorf("for %v and %v wanted %v, got %v", test.n, test.k, test.want, res)
		}
	}
}
