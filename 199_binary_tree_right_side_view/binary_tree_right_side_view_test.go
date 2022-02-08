package binary_tree_right_side_view

import (
	"reflect"
	"testing"
)

type Test struct {
	in  *TreeNode
	out []int
}

var tree1 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:   2,
		Right: &TreeNode{Val: 5},
	},
	Right: &TreeNode{
		Val:   3,
		Right: &TreeNode{Val: 4},
	},
}

var tree2 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:   2,
		Right: &TreeNode{Val: 5},
	},
	Right: &TreeNode{
		Val: 3,
	},
}

var tests = []Test{
	{&TreeNode{Val: 1}, []int{1}},
	{nil, []int{}},
	{tree1, []int{1, 3, 4}},
	{tree2, []int{1, 3, 5}},
}

func TestRightSideView(t *testing.T) {
	for _, test := range tests {
		if res := rightSideView(test.in); !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
