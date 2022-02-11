package merge_k_sorted_lists

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	queue := CreateQueue()

	for _, list := range lists {
		curr := list
		for curr != nil {
			queue.Push(curr.Val)
			curr = curr.Next
		}
	}

	sentinel := &ListNode{Val: -1}
	curr := sentinel

	for !queue.Empty() {
		curr.Next = &ListNode{Val: queue.Pop()}
		curr = curr.Next
	}

	return sentinel.Next
}

func mergeKListsBF(lists []*ListNode) *ListNode {
	all := make([]int, 0, 1024)

	for _, list := range lists {
		curr := list
		for curr != nil {
			all = append(all, curr.Val)
			curr = curr.Next
		}
	}

	sort.Ints(all)
	sentinel := &ListNode{Val: -1}
	curr := sentinel

	for _, v := range all {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}

	return sentinel.Next
}

func mergeKListsIter(lists []*ListNode) *ListNode {
	sentinel := &ListNode{Val: -1}
	head := sentinel

	for {
		idx := next(lists)

		if idx == -1 {
			break
		}

		head.Next = lists[idx]
		lists[idx] = lists[idx].Next
		head = head.Next
		head.Next = nil
	}

	return sentinel.Next
}

func next(lists []*ListNode) int {
	min := -1

	for i := range lists {
		if lists[i] == nil {
			continue
		}

		if min == -1 {
			min = i
			continue
		}

		if lists[min].Val > lists[i].Val {
			min = i
		}
	}

	return min
}
