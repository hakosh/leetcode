package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: -1}
	list := head

	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		v1, v2 := 0, 0

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		carry = sum / 10
		sum = sum % 10

		list.Next = &ListNode{Val: sum}
		list = list.Next
	}

	return head.Next
}
