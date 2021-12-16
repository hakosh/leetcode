package symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val == right.Val {
		return isMirror(left.Right, right.Left) && isMirror(left.Left, right.Right)
	}

	return false
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isSymmetricBFS(root *TreeNode) bool {
	lQ := []*TreeNode{root.Left}
	rQ := []*TreeNode{root.Right}

	for {
		if len(lQ) == 0 && len(rQ) == 0 {
			break
		}

		if len(lQ) != len(rQ) {
			break
		}

		var lNext []*TreeNode
		var rNext []*TreeNode

		for i, lNode := range lQ {
			rNode := rQ[i]

			if lNode == nil && rNode == nil {
				continue
			}

			if lNode == nil || rNode == nil {
				return false
			}

			if lNode.Val != rNode.Val {
				return false
			}

			lNext = append(lNext, lNode.Left)
			rNext = append(rNext, rNode.Right)
			lNext = append(lNext, lNode.Right)
			rNext = append(rNext, rNode.Left)
		}

		lQ = lNext
		rQ = rNext
	}

	return true
}
