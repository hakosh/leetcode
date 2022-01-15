package construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"reflect"
	"testing"
)

type Test struct {
	in  []int
	pre []int
}

var tests = []Test{
	{[]int{9, 3, 15, 20, 7}, []int{3, 9, 20, 15, 7}},
}

func TestBuildTree(t *testing.T) {
	for _, test := range tests {
		res := preorderTraversal(buildTree(test.pre, test.in))

		if reflect.DeepEqual(res, test.pre) {
			t.Errorf("For %v and %v got %v", test.in, test.pre, res)
		}
	}
}

func BenchmarkBuildTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildTree(tests[0].in, tests[0].pre)
	}
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	list := []int{root.Val}
	list = append(list, preorderTraversal(root.Left)...)
	list = append(list, preorderTraversal(root.Right)...)

	return list
}
