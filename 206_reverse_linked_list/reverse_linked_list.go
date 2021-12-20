package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	last := head

	for head.Next != nil {
		node := head.Next

		head.Next = node.Next
		node.Next = last

		last = node
	}

	return last
}
