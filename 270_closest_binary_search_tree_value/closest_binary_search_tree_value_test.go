package closest_binary_search_tree_value

import "testing"

type Test struct {
	root   *TreeNode
	target float64
	out    int
}

var tree1 = &TreeNode{
	Val: 4,
	Left: &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	},
	Right: &TreeNode{Val: 5},
}

var tree2 = &TreeNode{
	Val:   1,
	Left:  nil,
	Right: &TreeNode{Val: 2},
}

var tests = []Test{
	{tree1, 3.714286, 4},
	{&TreeNode{Val: 1}, 4.428571, 1},
	{tree2, 3.428571, 2},
}

func TestClosesValue(t *testing.T) {
	for _, test := range tests {
		if res := closestValue(test.root, test.target); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.root, test.target, test.out, res)
		}
	}
}
