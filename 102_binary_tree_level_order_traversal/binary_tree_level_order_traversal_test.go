package binary_tree_level_order_traversal

import (
	"reflect"
	"testing"
)

type Test struct {
	in  *TreeNode
	out [][]int
}

var tree1 = &TreeNode{
	Val: 1,
	Right: &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 3},
	},
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

var tree3 = &TreeNode{
	Val:  3,
	Left: &TreeNode{Val: 9},
	Right: &TreeNode{
		Val:   20,
		Left:  &TreeNode{Val: 15},
		Right: &TreeNode{Val: 7},
	},
}

var cases = []Test{
	{in: tree1, out: [][]int{{1}, {2}, {3}}},
	{in: tree2, out: [][]int{{9}, {4, 20}, {1, 6, 15, 170}}},
	{in: tree3, out: [][]int{{3}, {9, 20}, {15, 7}}},
	{in: nil, out: [][]int{}},
}

func TestLevelOrder(t *testing.T) {
	for _, test := range cases {
		if res := levelOrder(test.in); !reflect.DeepEqual(test.out, res) {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.out, res)
		}
	}
}
