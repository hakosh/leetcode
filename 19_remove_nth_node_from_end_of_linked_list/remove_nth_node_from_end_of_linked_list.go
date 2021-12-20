package remove_nth_node_from_end_of_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func length(head *ListNode) int {
	i := 1

	for head.Next != nil {
		head = head.Next
		i++
	}

	return i
}

func removeNthFromEnd_2Pass(head *ListNode, n int) *ListNode {
	to := length(head) - n

	if to == 0 {
		return head.Next
	}

	node := head
	for i := 1; i < to; i++ {
		node = node.Next
	}

	if n == 1 {
		node.Next = nil
	} else {
		node.Next = node.Next.Next
	}

	return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	back, front := head, head

	for ; n > 0; n-- {
		front = front.Next
	}

	if front == nil {
		return back.Next
	}

	for front.Next != nil {
		back = back.Next
		front = front.Next
	}

	if back.Next != nil {
		back.Next = back.Next.Next
	}

	return head
}
