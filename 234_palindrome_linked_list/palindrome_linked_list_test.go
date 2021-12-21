package palindrome_linked_list

import (
	"testing"
)

type Test struct {
	in  *ListNode
	out bool
}

var list1 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{Val: 1},
		},
	},
}

var list2 = &ListNode{
	Val: 5,
}

var list3 = &ListNode{
	Val: 8,
	Next: &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	},
}

var list4 = &ListNode{
	Val: 7,
	Next: &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val:  7,
				Next: &ListNode{Val: 7},
			},
		},
	},
}

var tests = []Test{
	{in: list1, out: true},
	{in: list2, out: true},
	{in: list3, out: false},
	{in: list4, out: true},
	{in: nil, out: true},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range tests {
		res := isPalindrome(test.in)

		if res != test.out {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.out, res)
		}
	}
}
