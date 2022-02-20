package sum_of_root_to_leaf_binary_numbers

import "testing"

type Test struct {
	root *TreeNode
	out  int
}

var tree1 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:   0,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 1},
	},
	Right: &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 1},
	},
}

var tests = []Test{
	{tree1, 22},
}

func TestSumRootToLeaf(t *testing.T) {
	for _, test := range tests {
		if res := sumRootToLeaf(test.root); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.root, test.out, res)
		}
	}
}
