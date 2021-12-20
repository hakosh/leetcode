package remove_nth_node_from_end_of_linked_list

import (
	"reflect"
	"testing"
)

type Test struct {
	in  *ListNode
	out bool
}

func traverse(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	return append([]int{head.Val}, traverse(head.Next)...)
}

var node1 = &ListNode{
	Val:  1,
	Next: node2,
}

var node2 = &ListNode{
	Val:  2,
	Next: node3,
}

var node3 = &ListNode{
	Val:  3,
	Next: node4,
}

var node4 = &ListNode{
	Val:  4,
	Next: nil,
}

var node5 = &ListNode{
	Val:  5,
	Next: node6,
}

var node6 = &ListNode{
	Val:  6,
	Next: nil,
}

func TestRemoveNthFromEnd(t *testing.T) {
	res := removeNthFromEnd(node1, 1)
	removed := traverse(res)

	if !reflect.DeepEqual(removed, []int{1, 2, 3}) {
		t.Errorf("Expected %v, got %v\n", []int{1, 2, 3}, removed)
	}

	removeNthFromEnd(node4, 1)

	res2 := removeNthFromEnd(node5, 2)
	removed = traverse(res2)

	if !reflect.DeepEqual(removed, []int{6}) {
		t.Errorf("Expected %v, got %v\n", []int{6}, removed)
	}
}
