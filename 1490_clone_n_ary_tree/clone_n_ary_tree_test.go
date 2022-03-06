package clone_n_ary_tree

import "testing"

type Test struct {
	root *Node
}

var tree1 = &Node{
	Val: 1,
	Children: []*Node{
		{Val: 2},
		{Val: 3},
	},
}

var tests = []Test{
	{tree1},
}

func TestCloneTree(t *testing.T) {
	for _, test := range tests {
		clone := cloneTree(test.root)
		traverse(t, test.root, clone)
	}
}

func traverse(t *testing.T, base, clone *Node) {
	if base == nil {
		return
	}

	if base.Val != clone.Val {
		t.Errorf("Not matching node: %d and %d", base.Val, clone.Val)
	} else {
		for i := range base.Children {
			traverse(t, base.Children[i], clone.Children[i])
		}
	}
}
