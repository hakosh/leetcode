package closest_binary_search_tree_value

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var closest int

func closestValue(root *TreeNode, target float64) int {
	closest = root.Val
	search(root, target)
	return closest
}

func search(root *TreeNode, target float64) {
	if root == nil {
		return
	}

	v := float64(root.Val)

	if abs(v-target) < abs(float64(closest)-target) {
		closest = root.Val
	}

	if target < v {
		search(root.Left, target)
	} else {
		search(root.Right, target)
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}

	return x
}
