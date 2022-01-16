package lowest_common_ancestor

import "testing"

type Test struct {
	tree  *TreeNode
	left  *TreeNode
	right *TreeNode
	lca   *TreeNode
}

var tree = &TreeNode{
	Val: 3,
	Left: &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 6},
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 7},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 10},
			},
		},
	},
	Right: &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 8},
	},
}

var node5 = tree.Left
var node4 = tree.Left.Right.Right

var tests = []Test{
	{tree, node5, node4, node5},
	{tree, node9, node10, node4},
}

func TestLowestCommonAncestor(t *testing.T) {
	for _, test := range tests {
		if res := lowestCommonAncestor(test.tree, test.left, test.right); res != test.lca {
			t.Errorf("For %v expected %v, got %v", test.tree, test.lca, res)
		}
	}
}

var node9 = tree.Left.Right.Right.Left
var node10 = tree.Left.Right.Right.Right

func BenchmarkLowestCommonAncestor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := tests[0]
		lowestCommonAncestor(t.tree, t.left, t.right)
	}
}
