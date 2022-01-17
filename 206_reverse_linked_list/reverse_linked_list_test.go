package reverse_linked_list

import (
	"reflect"
	"testing"
)

type Test struct {
	in  *ListNode
	out []int
}

var list1 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{Val: 4},
		},
	},
}

var list2 = &ListNode{
	Val: 5,
}

var list3 = &ListNode{
	Val:  8,
	Next: &ListNode{Val: 9},
}

var tests = []Test{
	{in: list1, out: []int{4, 3, 2, 1}},
	{in: list2, out: []int{5}},
	{in: list3, out: []int{9, 8}},
	{in: nil, out: []int{}},
}

func TestReverseList(t *testing.T) {
	for _, test := range tests {
		res := reverseList(test.in)

		if !reflect.DeepEqual(traverse(res), test.out) {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.out, traverse(res))
		}
	}
}

func traverse(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	return append([]int{head.Val}, traverse(head.Next)...)
}
