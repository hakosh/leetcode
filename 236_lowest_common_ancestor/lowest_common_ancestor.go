package lowest_common_ancestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var lca *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	lca = nil
	search(root, p, q)

	return lca
}

func search(root, p, q *TreeNode) bool {
	if root == nil {
		return false
	}

	v := 0

	if root == q || root == p {
		v++
	}

	if search(root.Left, p, q) {
		v++
	}

	if search(root.Right, p, q) {
		v++
	}

	if v >= 2 {
		lca = root
		return true
	} else if v > 0 {
		return true
	} else {
		return false
	}
}
