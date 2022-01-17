package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListIterative(head *ListNode) *ListNode {
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

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tail := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return tail
}
