package convert_sorted_array_to_binary_search_tree

import (
	"testing"
)

type Test struct {
	in  []int
	out *TreeNode
}

var tree1 = &TreeNode{
	Val: 0,
	Left: &TreeNode{
		Val:   -10,
		Right: &TreeNode{Val: -3},
	},
	Right: &TreeNode{
		Val:   5,
		Right: &TreeNode{Val: 9},
	},
}

var tree2 = &TreeNode{
	Val:   1,
	Right: &TreeNode{Val: 3},
}

var tests = []Test{
	{[]int{-10, -3, 0, 5, 9}, tree1},
	{[]int{1, 3}, tree2},
}

func TestSortedArrayToBST2(t *testing.T) {
	test := tests[1]
	tree := sortedArrayToBST(test.in)

	if tree.Val != 3 {
		t.Errorf("Root should be 1")
	}

	if tree.Left.Val != 1 {
		t.Errorf("Right child should be 3")
	}
}
