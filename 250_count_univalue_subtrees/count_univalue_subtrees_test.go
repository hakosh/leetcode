package count_univalue_subtrees

import "testing"

type Test struct {
	in  *TreeNode
	out int
}

var tree1 = &TreeNode{
	Val: 5,
	Left: &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 5},
		Right: &TreeNode{Val: 5},
	},
	Right: &TreeNode{
		Val:   5,
		Right: &TreeNode{Val: 5},
	},
}

var tree2 = &TreeNode{
	Val: 5,
	Left: &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 5},
		Right: &TreeNode{Val: 5},
	},
	Right: &TreeNode{
		Val:   5,
		Right: &TreeNode{Val: 5},
	},
}

var tests = []Test{
	{tree1, 4},
	{tree2, 6},
	{&TreeNode{Val: 1}, 1},
	{nil, 0},
}

func TestCountUnivalueSubtrees(t *testing.T) {
	for _, test := range tests {
		if res := countUnivalSubtrees(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}

func BenchmarkCountUnivalueSubtrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countUnivalSubtrees(tree1)
	}
}
