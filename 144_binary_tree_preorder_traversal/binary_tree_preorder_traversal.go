package binary_tree_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversalRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	list := []int{root.Val}
	list = append(list, preorderTraversalRecursive(root.Left)...)
	list = append(list, preorderTraversalRecursive(root.Right)...)

	return list
}

func preorderTraversal(root *TreeNode) []int {
	list := make([]int, 0)

	curr := root
	var stack []*TreeNode

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			list = append(list, curr.Val)
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		curr = curr.Right
	}

	return list
}
