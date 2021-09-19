package __palindrome_number

import "testing"

type testCase struct {
	number int
	result bool
}

var testCases = []testCase{
	{121, true},
	{-121, false},
	{0, true},
	{1, true},
	{9876789, true},
	{432234, true},
	{789789, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range testCases {
		if res := IsPalindrome(test.number); res != test.result {
			t.Errorf("%d: %t, got %t\n", test.number, test.result, res)
		}
	}
}

func TestIsPalindrome2(t *testing.T) {
	for _, test := range testCases {
		if res := IsPalindrome2(test.number); res != test.result {
			t.Errorf("%d: %t, got %t\n", test.number, test.result, res)
		}
	}
}
