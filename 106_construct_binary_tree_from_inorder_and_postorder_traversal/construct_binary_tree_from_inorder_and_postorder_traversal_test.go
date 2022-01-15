package construct_binary_tree_from_inorder_and_postorder_traversal

import (
	"reflect"
	"testing"
)

type Test struct {
	in   []int
	post []int
}

var tests = []Test{
	{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}},
}

func TestBuildTree(t *testing.T) {
	for _, test := range tests {
		res := inorderTraversal(buildTree(test.in, test.post))

		if reflect.DeepEqual(res, test.in) {
			t.Errorf("For %v and %v got %v", test.in, test.post, res)
		}
	}
}

func BenchmarkBuildTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildTree(tests[0].in, tests[0].post)
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
