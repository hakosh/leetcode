package binary_tree_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	list := []int{root.Val}
	list = append(list, preorderTraversal(root.Left)...)
	list = append(list, preorderTraversal(root.Right)...)

	return list
}
