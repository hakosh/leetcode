package binary_tree_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var list []int
	list = append(list, postorderTraversal(root.Left)...)
	list = append(list, postorderTraversal(root.Right)...)
	list = append(list, root.Val)

	return list
}
