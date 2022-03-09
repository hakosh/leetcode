package remove_duplicates_from_sorted_list_ii

import (
	"reflect"
	"testing"
)

type Test struct {
	head *ListNode
	want []int
}

var list1 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	},
}

var list2 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	},
}

var list3 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{Val: 4},
				},
			},
		},
	},
}

var list4 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{Val: 2},
		},
	},
}

var tests = []Test{
	{list1, []int{2, 3}},
	{list2, []int{1, 3, 4}},
	{list3, []int{1, 4}},
	{list4, []int{2}},
}

func TestDeleteDuplicates(t *testing.T) {
	for _, test := range tests {
		res := traverse(deleteDuplicates(test.head))

		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("For %v expected %v, got %v", test.head, test.want, res)
		}
	}
}

func traverse(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	return append([]int{head.Val}, traverse(head.Next)...)
}
