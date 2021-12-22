package copy_list_with_random_pointer

import (
	"reflect"
	"testing"
)

type Test struct {
	head *Node
	out  []int
}

var n1 = &Node{
	Val:    7,
	Next:   n2,
	Random: nil,
}

var n2 = &Node{
	Val:    13,
	Next:   n3,
	Random: nil,
}

var n3 = &Node{
	Val:    11,
	Next:   n4,
	Random: nil,
}

var n4 = &Node{
	Val:    10,
	Next:   n5,
	Random: nil,
}

var n5 = &Node{
	Val:    1,
	Next:   nil,
	Random: nil,
}

var tests = []Test{
	{n1, []int{7, 13, 11, 10, 1}},
	{nil, []int{}},
}

func TestCopyRandomList(t *testing.T) {
	for _, test := range tests {
		res := copyRandomList(test.head)

		if !reflect.DeepEqual(traverse(res), test.out) {
			t.Errorf("For %v wanted %v, got %v\n", test.head, test.out, traverse(res))
		}
	}
}

func traverse(head *Node) []int {
	if head == nil {
		return []int{}
	}

	return append([]int{head.Val}, traverse(head.Next)...)
}
