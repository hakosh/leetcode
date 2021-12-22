package insert_into_a_sorted_circular_linked_list

import (
	"reflect"
	"testing"
)

type Test struct {
	list *Node
	x    int
	out  []int
}

var l1 = &Node{
	Val: 3,
	Next: &Node{
		Val: 4,
		Next: &Node{
			Val:  1,
			Next: nil,
		},
	},
}

var l2 = &Node{Val: 5}

var l3 = &Node{
	Val: 3,
	Next: &Node{
		Val: 3,
		Next: &Node{
			Val:  3,
			Next: nil,
		},
	},
}

var l4 = &Node{
	Val: 3,
	Next: &Node{
		Val: 5,
		Next: &Node{
			Val:  1,
			Next: nil,
		},
	},
}

var tests = []Test{
	{list: l1, x: 2, out: []int{3, 4, 1, 2}},
	{list: l2, x: 2, out: []int{5, 2}},
	{list: l3, x: 0, out: []int{3, 3, 3, 0}},
	{list: l4, x: 0, out: []int{3, 5, 0, 1}},
	{list: nil, x: 1, out: []int{1}},
}

func TestFlatten(t *testing.T) {
	l1.Next.Next.Next = l1
	l2.Next = l2
	l3.Next.Next.Next = l3
	l4.Next.Next.Next = l4

	for _, test := range tests {
		res := insert(test.list, test.x)

		if !reflect.DeepEqual(traverse(res, len(test.out)), test.out) {
			t.Errorf("For %v wanted %v, got %v\n", test.list, test.out, traverse(res, len(test.out)))
		}
	}
}

func traverse(head *Node, until int) []int {
	if head == nil || until == 0 {
		return []int{}
	}

	return append([]int{head.Val}, traverse(head.Next, until-1)...)
}
