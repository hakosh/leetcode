package add_two_numbers

import (
	"reflect"
	"testing"
)

type Test struct {
	l1  *ListNode
	l2  *ListNode
	out []int
}

var l1 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	},
}

var l2 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	},
}

var l3 = &ListNode{
	Val: 2,
	Next: &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	},
}

var l4 = &ListNode{
	Val: 5,
	Next: &ListNode{
		Val: 6,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	},
}

var tests = []Test{
	{l1, l2, []int{2, 5, 8}},
	{l3, l4, []int{7, 0, 8}},
	{&ListNode{Val: 0}, &ListNode{Val: 0}, []int{0}},
}

func TestAddTwoNumbers(t *testing.T) {
	for _, test := range tests {
		res := addTwoNumbers(test.l1, test.l2)

		if !reflect.DeepEqual(traverse(res), test.out) {
			t.Errorf("For %v %v wanted %v, got %v\n", test.l1, test.l2, test.out, traverse(res))
		}
	}
}

func traverse(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	return append([]int{head.Val}, traverse(head.Next)...)
}
