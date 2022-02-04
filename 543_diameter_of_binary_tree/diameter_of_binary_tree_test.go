package diameter_of_binary_tree

import "testing"

type Test struct {
	tree *TreeNode
	out  int
}

var tree1 = &TreeNode{
	Val:  1,
	Left: &TreeNode{Val: 2},
}

var tree2 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   6,
			Right: &TreeNode{Val: 7},
		},
	},
	Right: &TreeNode{Val: 3},
}

var tests = []Test{
	{&TreeNode{Val: 1}, 0},
	{tree1, 1},
	{tree2, 4},
}

func TestDiameterOfBinaryTree(t *testing.T) {
	for _, test := range tests {
		if res := diameterOfBinaryTree(test.tree); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.tree, test.out, res)
		}
	}
}
