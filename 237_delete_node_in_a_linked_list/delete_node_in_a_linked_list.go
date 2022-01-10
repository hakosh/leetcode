package delete_node_in_a_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	for node.Next.Next != nil {
		node.Val = node.Next.Val
		node = node.Next
	}

	node.Val = node.Next.Val
	node.Next = nil
}
