package diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var diameter int

func diameterOfBinaryTree(root *TreeNode) int {
	diameter = 0
	check(root)
	return diameter
}

func check(node *TreeNode) int {
	if node == nil {
		return 0
	}

	var depthLeft, depthRight int

	if node.Left != nil {
		depthLeft = check(node.Left) + 1
	}

	if node.Right != nil {
		depthRight = check(node.Right) + 1
	}

	diameter = max(depthLeft+depthRight, diameter)

	return max(depthLeft, depthRight)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
