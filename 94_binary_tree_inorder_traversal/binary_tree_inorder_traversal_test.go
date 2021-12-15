package binary_tree_postorder_traversal

import (
	"reflect"
	"testing"
)

type Test struct {
	in  *TreeNode
	out []int
}

var tree2 = &TreeNode{
	Val: 9,
	Left: &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 6},
	},
	Right: &TreeNode{
		Val:   20,
		Left:  &TreeNode{Val: 15},
		Right: &TreeNode{Val: 170},
	},
}

var cases = []Test{
	{in: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, out: []int{1, 3, 2}},
	{in: tree2, out: []int{1, 4, 6, 9, 15, 20, 170}},
	{in: nil, out: []int{}},
}

func TestInorderTraversal(t *testing.T) {
	for _, test := range cases {
		if res := inorderTraversal(test.in); !reflect.DeepEqual(test.out, res) {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.out, res)
		}
	}
}
