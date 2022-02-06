package reverse_linked_list_ii

import (
	"reflect"
	"testing"
)

type Test struct {
	in    *ListNode
	left  int
	right int
	out   []int
}

var list1 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: &ListNode{Val: 6},
				},
			},
		},
	},
}

var list2 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{Val: 4},
		},
	},
}

var list3 = &ListNode{Val: 5}

var list4 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{Val: 4},
		},
	},
}

var list5 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{Val: 4},
		},
	},
}

var tests = []Test{
	{list1, 2, 5, []int{1, 5, 4, 3, 2, 6}},
	{list2, 2, 3, []int{1, 3, 2, 4}},
	{list3, 1, 1, []int{5}},
	{list4, 1, 4, []int{4, 3, 2, 1}},
	{list5, 4, 4, []int{1, 2, 3, 4}},
}

func TestReverseBetween(t *testing.T) {
	for _, test := range tests {
		list := reverseBetween(test.in, test.left, test.right)

		if res := traverse(list); !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}

func traverse(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	return append([]int{head.Val}, traverse(head.Next)...)
}
