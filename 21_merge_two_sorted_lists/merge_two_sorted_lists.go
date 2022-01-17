package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoListsIterative(list1 *ListNode, list2 *ListNode) *ListNode {
	merged := &ListNode{Val: -1}
	mergedHead := merged

	for list1 != nil || list2 != nil {
		if list1 == nil {
			merged.Next = list2
			list2 = list2.Next
		} else if list2 == nil || list1.Val < list2.Val {
			merged.Next = list1
			list1 = list1.Next
		} else {
			merged.Next = list2
			list2 = list2.Next
		}

		merged = merged.Next
	}

	return mergedHead.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var head *ListNode

	if list1.Val < list2.Val {
		head = list1
		head.Next = mergeTwoLists(list1.Next, list2)
	} else {
		head = list2
		head.Next = mergeTwoLists(list1, list2.Next)
	}

	return head
}
