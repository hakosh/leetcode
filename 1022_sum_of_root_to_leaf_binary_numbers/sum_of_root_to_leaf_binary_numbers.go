package sum_of_root_to_leaf_binary_numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum = 0

func sumRootToLeaf(root *TreeNode) int {
	sum = 0
	dfs(root, 0)
	return sum
}

func dfs(root *TreeNode, current int) {
	if root == nil {
		return
	}

	current = current << 1
	current += root.Val

	if root.Left == nil && root.Right == nil {
		sum += current
	} else {
		dfs(root.Left, current)
		dfs(root.Right, current)
	}
}
