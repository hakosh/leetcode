package path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var answer bool

func hasPathSum(root *TreeNode, targetSum int) bool {
	answer = false
	findSum(root, targetSum)
	return answer
}

func findSum(root *TreeNode, targetSum int) {
	if answer || root == nil {
		return
	}

	targetSum -= root.Val

	if root.Left == nil && root.Right == nil && targetSum == 0 {
		answer = true
		return
	}

	findSum(root.Left, targetSum)
	findSum(root.Right, targetSum)
}
