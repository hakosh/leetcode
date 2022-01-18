package unique_binary_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return gen(1, n)
}

func gen(from, to int) []*TreeNode {
	if from > to {
		return []*TreeNode{nil}
	}

	var trees []*TreeNode

	for i := from; i <= to; i++ {
		left := gen(from, i-1)
		right := gen(i+1, to)

		for _, lTree := range left {
			for _, rTree := range right {
				tree := &TreeNode{
					Val:   i,
					Left:  lTree,
					Right: rTree,
				}

				trees = append(trees, tree)
			}
		}
	}

	return trees
}
