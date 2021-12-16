package symmetric_tree

import (
	"testing"
)

type Test struct {
	in  *TreeNode
	out bool
}

var tree1 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:   2,
		Right: &TreeNode{Val: 3},
	},
	Right: &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 3},
	},
}

var tree2 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 4},
	},
	Right: &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 3},
	},
}

var tree3 = &TreeNode{
	Val:   3,
	Left:  &TreeNode{Val: 9},
	Right: &TreeNode{Val: 20},
}

var tree4 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 2},
	},
	Right: &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 2},
	},
}

var tree5 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:   2,
		Right: &TreeNode{Val: 3},
	},
	Right: &TreeNode{
		Val:   2,
		Right: &TreeNode{Val: 3},
	},
}

var cases = []Test{
	{in: tree1, out: true},
	{in: tree2, out: true},
	{in: tree3, out: false},
	{in: tree4, out: false},
	{in: tree5, out: false},
}

func TestIsSymmetric(t *testing.T) {
	for _, test := range cases {
		if res := isSymmetricBFS(test.in); test.out != res {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.out, res)
		}
	}
}
