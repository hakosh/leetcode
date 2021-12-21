package odd_even_linked_list

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
	{in: list1, out: []int{1, 3, 2}},
	{in: list2, out: []int{5}},
	{in: list3, out: []int{8, 8, 9, 9}},
	{in: list4, out: []int{7, 7, 7, 7}},
	{in: nil, out: []int{}},
}

func TestOddEvenList(t *testing.T) {
	for _, test := range tests {
		res := oddEvenList(test.in)

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
