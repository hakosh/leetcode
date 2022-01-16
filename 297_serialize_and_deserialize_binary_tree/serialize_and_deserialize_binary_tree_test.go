package serialize_and_deserialize_binary_tree

import (
	"fmt"
	"reflect"
	"testing"
)

var tree = &TreeNode{
	Val: 3,
	Left: &TreeNode{
		Val:  -5,
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

func TestCodec(t *testing.T) {
	c := Constructor()

	tree2 := c.deserialize(c.serialize(tree))

	if !reflect.DeepEqual(inorderTraversal(tree), inorderTraversal(tree2)) {
		fmt.Println(c.serialize(tree))
		t.Errorf("The trees are not matching")
	}

	empty := c.deserialize(c.serialize(nil))

	if empty != nil {
		t.Errorf("Empty tree is not empty")
	}
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var list []int
	list = append(list, inorderTraversal(root.Left)...)
	list = append(list, root.Val)
	list = append(list, inorderTraversal(root.Right)...)

	return list
}

func BenchmarkSerialize(b *testing.B) {
	c := Constructor()

	for i := 0; i < b.N; i++ {
		c.serialize(tree)
	}
}

func BenchmarkDeserialize(b *testing.B) {
	c := Constructor()
	d := c.serialize(tree)

	for i := 0; i < b.N; i++ {
		c.deserialize(d)
	}
}
