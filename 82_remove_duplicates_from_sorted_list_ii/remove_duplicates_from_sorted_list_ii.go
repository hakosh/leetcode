package remove_duplicates_from_sorted_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	for head != nil && head.Next != nil && head.Val == head.Next.Val {
		val := head.Val

		for head != nil && head.Val == val {
			head = head.Next
		}
	}

	if head == nil {
		return head
	} else {
		head.Next = deleteDuplicates(head.Next)
		return head
	}
}
