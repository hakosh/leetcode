package count_complete_tree_nodes

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	height := getHeight(root)

	last := 0
	nodes := 0
	for i := 0; i < height; i++ {
		last = int(math.Pow(2, float64(i)))
		nodes += last
	}

	return nodes - last + findLast(root, last)
}

func getHeight(root *TreeNode) int {
	h := 0

	for root != nil {
		h++
		root = root.Left
	}

	return h
}

func findLast(root *TreeNode, total int) int {
	left, right := 1, total

	for left <= right {
		mid := (left + right) / 2

		if check(root, mid, total) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left - 1
}

func check(root *TreeNode, target, total int) bool {
	left, right := 1, total

	for left != right {
		mid := (left + right) / 2

		if target <= mid {
			root = root.Left
			right = mid
		} else {
			root = root.Right
			left = mid + 1
		}
	}

	return root != nil
}
