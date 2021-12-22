package flatten_a_multilevel_doubly_linked_list

import (
	"reflect"
	"testing"
)

type Test struct {
	list *Node
	out  []int
}

var n1 = &Node{
	Val:  1,
	Next: n2,
}

var n2 = &Node{
	Val:  2,
	Next: n3,
}

var n3 = &Node{
	Val:   3,
	Child: n1a,
}

var n1a = &Node{
	Val:  7,
	Next: n2a,
	Prev: nil,
}

var n2a = &Node{
	Val:   8,
	Next:  n3a,
	Child: n1b,
}

var n3a = &Node{
	Val:  9,
	Next: nil,
}

var n1b = &Node{
	Val:  11,
	Next: n2b,
	Prev: nil,
}

var n2b = &Node{
	Val:  12,
	Prev: nil, // n1b
	Next: nil,
}

var tests = []Test{
	{list: n1, out: []int{1, 2, 3, 7, 8, 11, 12, 9}},
	{list: nil, out: []int{}},
}

func TestFlatten(t *testing.T) {
	n2.Prev = n1
	n3.Prev = n2
	n2a.Prev = n1a
	n3a.Prev = n2a
	n2b.Prev = n1b

	for _, test := range tests {
		res := flatten(test.list)

		if !reflect.DeepEqual(traverse(res), test.out) {
			t.Errorf("For %v wanted %v, got %v\n", test.list, test.out, traverse(res))
		}
	}
}

func traverse(head *Node) []int {
	if head == nil {
		return []int{}
	}

	return append([]int{head.Val}, traverse(head.Next)...)
}
