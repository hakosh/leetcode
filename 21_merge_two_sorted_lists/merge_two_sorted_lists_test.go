package merge_two_sorted_lists

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

var tests = []Test{
	{l1, l2, []int{1, 1, 2, 3, 4, 4}},
	{nil, nil, []int{}},
	{nil, &ListNode{Val: 0}, []int{0}},
}

func TestMergeTwoLists(t *testing.T) {
	for _, test := range tests {
		res := mergeTwoLists(test.l1, test.l2)

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
