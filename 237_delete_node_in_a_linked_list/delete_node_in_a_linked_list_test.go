package delete_node_in_a_linked_list

import (
	"reflect"
	"testing"
)

type Test struct {
	list *ListNode
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

var tests = []Test{
	{list: l1, out: []int{1, 2, 4, 5}},
}

func TestDeleteNode(t *testing.T) {
	for _, test := range tests {
		deleteNode(l1.Next.Next)

		if !reflect.DeepEqual(traverse(l1), test.out) {
			t.Errorf("For %v wanted %v, got %v\n", test.list, test.out, traverse(l1))
		}
	}
}

func traverse(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	return append([]int{head.Val}, traverse(head.Next)...)
}
