package remove_linked_list_elements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	cur := head
	newHead := head
	var prev *ListNode

	for cur != nil {
		if cur.Val == val {
			if prev == nil {
				newHead = cur.Next
			} else {
				prev.Next = cur.Next
			}
		} else {
			prev = cur
		}

		cur = cur.Next
	}

	return newHead
}
