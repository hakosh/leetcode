package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListIterative(head *ListNode) *ListNode {
	var prev *ListNode
	var curr = head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
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
