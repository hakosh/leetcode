package populating_next_right_pointers_in_each_node_ii

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		tmp := make([]*Node, 0)

		for i, node := range queue {
			if i > 0 {
				queue[i-1].Next = node
			}

			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}

			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}

		queue = tmp
	}

	return root
}
