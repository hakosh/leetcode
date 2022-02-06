package reverse_linked_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	sentinel := &ListNode{Val: -1, Next: head}
	leftNode := sentinel
	for i := 1; i < left; i++ {
		leftNode = leftNode.Next
	}

	curr := leftNode.Next

	for i := 0; i < right-left; i++ {
		next := curr.Next

		curr.Next = next.Next
		next.Next = leftNode.Next
		leftNode.Next = next
	}

	return sentinel.Next
}
