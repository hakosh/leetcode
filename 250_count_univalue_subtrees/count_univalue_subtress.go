package count_univalue_subtrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countUnivalSubtrees(root *TreeNode) int {
	if root == nil {
		return 0
	}

	v := 0
	if isUnivalue(root) {
		v++
	}

	return v + countUnivalSubtrees(root.Left) + countUnivalSubtrees(root.Right)
}

func isUnivalue(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil {
		return true
	}

	univalue := true

	if root.Left != nil && root.Left.Val != root.Val {
		univalue = false
	}

	if root.Right != nil && root.Right.Val != root.Val {
		univalue = false
	}

	return univalue && isUnivalue(root.Left) && isUnivalue(root.Right)
}
