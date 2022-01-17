package search_in_a_binary_search_tree

import "testing"

func TestSearchBST(t *testing.T) {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}

	node3 := tree.Left.Right
	hit := searchBST(tree, 3)

	if hit != node3 {
		t.Errorf("Node 3 not found, instead %v", hit)
	}
}
