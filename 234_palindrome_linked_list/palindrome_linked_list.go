package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	tail := midpoint(head)
	tail = reverseList(tail)

	// Compare lists
	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}

		head, tail = head.Next, tail.Next
	}

	return true
}

func midpoint(head *ListNode) *ListNode {
	slow := head
	fast := head
	prev := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}

	if fast != nil {
		mid := slow.Next
		slow.Next = nil
		return mid
	} else {
		prev.Next = nil
		return slow
	}
}

func reverseList(head *ListNode) *ListNode {
	last := head

	for head.Next != nil {
		node := head.Next

		head.Next = node.Next
		node.Next = last

		last = node
	}

	return last
}
