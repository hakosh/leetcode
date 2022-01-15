package construct_binary_tree_from_inorder_and_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(inorder) == 0 {
		return nil
	}

	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}

	if len(postorder) == 1 {
		return &TreeNode{Val: postorder[0]}
	}

	idxLeft := -1
	idxRight := -1

	for i := len(postorder) - 1; i >= 0; i-- {
		for j, v := range inorder {
			if v == postorder[i] {
				idxLeft = j
				idxRight = i
				i = -1
				break
			}
		}
	}

	return &TreeNode{
		Val:   inorder[idxLeft],
		Left:  buildTree(inorder[:idxLeft], postorder[:idxRight]),
		Right: buildTree(inorder[idxLeft+1:], postorder),
	}
}
