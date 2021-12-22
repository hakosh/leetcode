package insert_into_a_sorted_circular_linked_list

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	node := &Node{Val: x}

	if aNode == nil {
		node.Next = node
		return node
	}

	tail := findLast(aNode)
	head := tail.Next

	if x >= tail.Val {
		node.Next = tail.Next
		tail.Next = node
	} else if x <= head.Val {
		node.Next = head
		tail.Next = node
	} else {
		for x > head.Next.Val {
			head = head.Next
		}

		node.Next = head.Next
		head.Next = node
	}

	return aNode
}

func findLast(node *Node) *Node {
	head := node

	for node.Next.Val >= node.Val && node.Next != head {
		node = node.Next
	}

	return node
}
