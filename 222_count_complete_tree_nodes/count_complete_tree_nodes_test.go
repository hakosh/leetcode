package count_complete_tree_nodes

import "testing"

type Test struct {
	in  *TreeNode
	out int
}

var tree1 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 5},
	},
	Right: &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 6},
	},
}

var tree2 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 5},
	},
	Right: &TreeNode{
		Val: 3,
	},
}

var tree3 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 8},
		},
		Right: &TreeNode{Val: 5},
	},
	Right: &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 6},
		Right: &TreeNode{Val: 7},
	},
}

var tests = []Test{
	{&TreeNode{Val: 1}, 1},
	{&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, 2},
	{&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, 3},
	{tree1, 6},
	{tree2, 5},
	{tree3, 8},
	{nil, 0},
}

func TestCountNodes(t *testing.T) {
	for _, test := range tests {
		if res := countNodes(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
