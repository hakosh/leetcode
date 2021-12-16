package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Breadth First Search
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var nodes [][]int
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		var level []int
		var next []*TreeNode

		for _, node := range queue {
			level = append(level, node.Val)

			if node.Left != nil {
				next = append(next, node.Left)
			}

			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		nodes = append(nodes, level)
		queue = next
	}

	return nodes
}
