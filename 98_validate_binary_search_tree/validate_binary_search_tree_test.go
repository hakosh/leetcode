package validate_binary_search_tree

import (
	"math"
	"testing"
)

type Test struct {
	in  *TreeNode
	out bool
}

var tree1 = &TreeNode{
	Val:  5,
	Left: &TreeNode{Val: 1},
	Right: &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 6},
	},
}

var tree2 = &TreeNode{
	Val: 5,
	Left: &TreeNode{
		Val: 4,
	},
	Right: &TreeNode{
		Val:   6,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 7},
	},
}

var tree3 = &TreeNode{
	Val:   2,
	Left:  &TreeNode{Val: 1},
	Right: &TreeNode{Val: 3},
}

var tests = []Test{
	{tree1, false},
	{tree2, false},
	{tree3, true},
	{&TreeNode{Val: 1}, true},
	{&TreeNode{Val: math.MaxInt32}, true},
}

func TestIsValidBST(t *testing.T) {
	for _, test := range tests {
		if res := isValidBST(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
