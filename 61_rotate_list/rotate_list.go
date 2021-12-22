package rotate_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func length(head *ListNode) int {
	i := 0

	for head != nil {
		head = head.Next
		i++
	}

	return i
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	l := length(head)
	k = k % l

	if k == 0 {
		return head
	}

	node := head

	var newHead *ListNode
	var newTail *ListNode

	for i := 1; i <= l; i++ {
		if i == l-k {
			newHead = node.Next
			newTail = node
		} else if i == l {
			node.Next = head
			newTail.Next = nil
		}

		node = node.Next
	}

	return newHead
}
