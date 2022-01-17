package swap_nodes_in_pairs

import (
	"reflect"
	"testing"
)

type Test struct {
	in  *ListNode
	out []int
}

var l1 = &ListNode{
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
	{l1, []int{2, 1, 4, 3}},
}

func TestDecodeString(t *testing.T) {
	for _, test := range tests {
		l := swapPairs(test.in)
		if !reflect.DeepEqual(traverse(l), test.out) {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, traverse(l))
		}
	}
}

func traverse(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	return append([]int{head.Val}, traverse(head.Next)...)
}
