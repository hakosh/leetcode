package remove_linked_list_elements

import (
	"reflect"
	"testing"
)

type Test struct {
	in  *ListNode
	val int
	out []int
}

var list1 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  3,
			Next: nil,
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
			Val: 7,
			Next: &ListNode{
				Val:  7,
				Next: nil,
			},
		},
	},
}

var tests = []Test{
	{in: list1, val: 2, out: []int{1, 3}},
	{in: list2, val: 5, out: []int{}},
	{in: list3, val: 8, out: []int{9, 9}},
	{in: list4, val: 7, out: []int{}},
	{in: nil, val: 0, out: []int{}},
}

func TestRemoveElements(t *testing.T) {
	for _, test := range tests {
		res := removeElements(test.in, test.val)

		if !reflect.DeepEqual(traverse(res), test.out) {
			t.Errorf("For %v and %v wanted %v, got %v\n", test.in, test.val, test.out, traverse(res))
		}
	}
}

func traverse(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	return append([]int{head.Val}, traverse(head.Next)...)
}
