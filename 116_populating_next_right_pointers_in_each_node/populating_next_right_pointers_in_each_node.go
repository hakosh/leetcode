package populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}

	link(root.Left, root.Right)

	return root
}

func link(src, dst *Node) {
	src.Next = dst

	if src.Right != nil {
		link(src.Right, dst.Left)
		link(src.Left, src.Right)
		link(dst.Left, dst.Right)
	}
}
