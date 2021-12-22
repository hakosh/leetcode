package rotate_list

import (
	"reflect"
	"testing"
)

type Test struct {
	list *ListNode
	k    int
	out  []int
}

var l1 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: &ListNode{Val: 5},
			},
		},
	},
}

var l2 = &ListNode{
	Val: 0,
	Next: &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	},
}

var l3 = &ListNode{Val: 9}

var l4 = &ListNode{
	Val: 7,
	Next: &ListNode{
		Val: 8,
	},
}

var l5 = &ListNode{
	Val:  1,
	Next: &ListNode{Val: 2},
}

var tests = []Test{
	{l1, 2, []int{4, 5, 1, 2, 3}},
	{l2, 4, []int{2, 0, 1}},
	{l3, 4, []int{9}},
	{l4, 97, []int{8, 7}},
	{l5, 2, []int{1, 2}},
	{nil, 2, []int{}},
}

func TestAddTwoNumbers(t *testing.T) {
	for _, test := range tests {
		res := rotateRight(test.list, test.k)

		if !reflect.DeepEqual(traverse(res), test.out) {
			t.Errorf("For %v %v wanted %v, got %v\n", test.list, test.k, test.out, traverse(res))
		}
	}
}

func traverse(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	return append([]int{head.Val}, traverse(head.Next)...)
}
