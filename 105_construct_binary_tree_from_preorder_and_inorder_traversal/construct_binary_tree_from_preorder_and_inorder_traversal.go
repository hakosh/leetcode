package construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 || len(inorder) == 0 {
		return nil
	}

	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}

	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	idxIn := -1
	idxPre := -1

	for i := 0; i < len(preorder); i++ {
		for j, m := range inorder {
			if m == preorder[i] {
				idxIn = j
				idxPre = i
				i = len(preorder)
				break
			}
		}
	}

	return &TreeNode{
		Val:   inorder[idxIn],
		Left:  buildTree(preorder[idxPre:], inorder[:idxIn]),
		Right: buildTree(preorder, inorder[idxIn+1:]),
	}
}
