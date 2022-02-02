package validate_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return check(root, nil, nil)
}

func check(root *TreeNode, low, high *int) bool {
	if root == nil {
		return true
	}

	if (low != nil && root.Val <= *low) || (high != nil && root.Val >= *high) {
		return false
	}

	return check(root.Right, &root.Val, high) && check(root.Left, low, &root.Val)
}
