package binary_tree_preorder_traversal

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
	{in: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, out: []int{1, 2, 3}},
	{in: tree2, out: []int{9, 4, 1, 6, 20, 15, 170}},
	{in: nil, out: []int{}},
}

func TestThirdMax(t *testing.T) {
	for _, test := range cases {
		if res := preorderTraversal(test.in); !reflect.DeepEqual(test.out, res) {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.out, res)
		}
	}
}
