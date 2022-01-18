package unique_binary_trees

import "testing"

type Test struct {
	in  int
	out []*TreeNode
}

var t1 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 2},
	},
}

var tests = []Test{
	{3, []*TreeNode{t1, nil, nil, nil, nil}},
}

func TestGenerateTrees(t *testing.T) {
	for _, test := range tests {
		trees := generateTrees(test.in)

		if len(trees) != len(test.out) {
			t.Errorf("Expected %v trees, got %v", len(test.out), len(trees))
		}
	}
}
