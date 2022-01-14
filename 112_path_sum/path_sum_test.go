package path_sum

import "testing"

type Test struct {
	in     *TreeNode
	target int
	out    bool
}

var tree1 = &TreeNode{
	Val: 5,
	Left: &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   11,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 2},
		},
	},
	Right: &TreeNode{
		Val:  8,
		Left: &TreeNode{Val: 13},
		Right: &TreeNode{
			Val:   4,
			Right: &TreeNode{Val: 1},
		},
	},
}

var tree2 = &TreeNode{
	Val:   1,
	Left:  &TreeNode{Val: 2},
	Right: &TreeNode{Val: 3},
}

var tests = []Test{
	{tree1, 22, true},
	{tree2, 5, false},
}

func TestHasPathSum(t *testing.T) {
	for _, test := range tests {
		if res := hasPathSum(test.in, test.target); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.in, test.target, test.out, res)
		}
	}
}
