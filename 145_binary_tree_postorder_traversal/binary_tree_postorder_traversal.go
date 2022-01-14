package binary_tree_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversalRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var list []int
	list = append(list, postorderTraversalRecursive(root.Left)...)
	list = append(list, postorderTraversalRecursive(root.Right)...)
	list = append(list, root.Val)

	return list
}

func postorderTraversal(root *TreeNode) []int {
	list := make([]int, 0)

	curr := root
	stack := make([]*TreeNode, 0)

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)

			if curr.Left != nil {
				curr = curr.Left
			} else {
				curr = curr.Right
			}
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		list = append(list, curr.Val)

		if len(stack) != 0 && curr == stack[len(stack)-1].Left && stack[len(stack)-1].Right != nil {
			curr = stack[len(stack)-1].Right
		} else {
			curr = nil
		}
	}

	return list
}
