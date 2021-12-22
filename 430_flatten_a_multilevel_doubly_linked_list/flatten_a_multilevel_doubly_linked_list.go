package flatten_a_multilevel_doubly_linked_list

type Node struct {
	Val   int
	Next  *Node
	Prev  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	head := root

	for root != nil {
		if root.Child == nil {
			root = root.Next
		} else {
			sub := flatten(root.Child)
			subTail := tail(sub)
			root.Child = nil

			subTail.Next = root.Next
			if root.Next != nil {
				root.Next.Prev = subTail
			}

			root.Next = sub
			sub.Prev = root

			root = subTail.Next
		}
	}

	return head
}

func tail(node *Node) *Node {
	for node.Next != nil {
		node = node.Next
	}

	return node
}
