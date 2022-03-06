package clone_n_ary_tree

type Node struct {
	Val      int
	Children []*Node
}

func cloneTree(root *Node) *Node {
	if root == nil {
		return root
	}

	node := &Node{Val: root.Val}

	for i := range root.Children {
		node.Children = append(node.Children, cloneTree(root.Children[i]))
	}

	return node
}
