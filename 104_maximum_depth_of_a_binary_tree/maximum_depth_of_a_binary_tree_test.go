package maximum_depth_of_a_binary_tree

import (
	"testing"
)

type Test struct {
	in  *TreeNode
	out int
}

var tree1 = &TreeNode{
	Val: 1,
	Right: &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 3},
	},
}

var tree2 = &TreeNode{
	Val: 9,
	Left: &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 6},
	},
	Right: &TreeNode{
		Val:   20,
		Left:  &TreeNode{Val: 15},
		Right: &TreeNode{Val: 170},
	},
}

var tree3 = &TreeNode{
	Val:   3,
	Left:  &TreeNode{Val: 9},
	Right: &TreeNode{Val: 20},
}

var cases = []Test{
	{in: tree1, out: 3},
	{in: tree2, out: 3},
	{in: tree3, out: 2},
	{in: nil, out: 0},
}

func TestLevelOrder(t *testing.T) {
	for _, test := range cases {
		if res := maxDepth(test.in); test.out != res {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.out, res)
		}
	}
}
