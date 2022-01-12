package binary_tree_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var list []int
	list = append(list, inorderTraversal(root.Left)...)
	list = append(list, root.Val)
	list = append(list, inorderTraversal(root.Right)...)

	return list
}

func inorderTraversalStack(root *TreeNode) []int {
	var values []int

	if root == nil {
		return []int{}
	}

	curr := root
	var stack []*TreeNode

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		values = append(values, curr.Val)

		curr = curr.Right
	}

	return values
}
