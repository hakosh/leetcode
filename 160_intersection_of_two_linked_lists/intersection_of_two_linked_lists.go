package intersection_of_two_linked_lists

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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := length(headA)
	lenB := length(headB)

	var diff int
	var long, short *ListNode

	if lenA > lenB {
		long, short = headA, headB
		diff = lenA - lenB
	} else {
		long, short = headB, headA
		diff = lenB - lenA
	}

	for i := 0; i < diff; i++ {
		long = long.Next
	}

	for long != nil {
		if long == short {
			return long
		}

		long = long.Next
		short = short.Next
	}

	return nil
}
