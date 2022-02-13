package maximum_length_of_a_concatenated_string_with_unique_characters

import "testing"

type Test struct {
	arr []string
	out int
}

var tests = []Test{
	{[]string{"un", "iq", "ue"}, 4},
	{[]string{"un", "iq", "ue", "ujo"}, 5},
	{[]string{"un", "iq", "ue", "ujo", "x", "xbc"}, 8},
	{[]string{"cha", "r", "act", "ers"}, 6},
	{[]string{"abcdefghijklmnopqrstuvwxyz"}, 26},
}

func TestMaxLength(t *testing.T) {
	for _, test := range tests {
		if res := maxLength(test.arr); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.arr, test.out, res)
		}
	}
}
