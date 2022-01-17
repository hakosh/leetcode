package swap_nodes_in_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	from := head.Next.Next

	head.Next.Next = head
	head.Next = swapPairs(from)

	return next
}
