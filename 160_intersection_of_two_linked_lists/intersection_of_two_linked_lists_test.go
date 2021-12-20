package intersection_of_two_linked_lists

import "testing"

type Test struct {
	head1 *ListNode
	head2 *ListNode
	out   *ListNode
}

var noCycle = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	},
}

var node1 = &ListNode{
	Val:  1,
	Next: node2,
}

var node2 = &ListNode{
	Val:  2,
	Next: node3,
}

var node3 = &ListNode{
	Val:  3,
	Next: node4,
}

var node4 = &ListNode{
	Val:  4,
	Next: nil,
}

var node1a = &ListNode{
	Val:  9,
	Next: node2a,
}

var node2a = &ListNode{
	Val:  8,
	Next: nil,
}

var node1b = &ListNode{
	Val:  0,
	Next: node1,
}

var tests = []Test{
	{head1: node1, head2: node1a, out: node3},
	{head1: node1b, head2: node1, out: node1},
	{head1: noCycle, head2: node1, out: nil},
}

func TestDetectCycle(t *testing.T) {
	node2a.Next = node3

	for _, test := range tests {
		res := getIntersectionNode(test.head1, test.head2)

		if res != test.out {
			t.Errorf("For %v and %v wanted %v, got %v\n", test.head1, test.head2, test.out, res)
		}
	}
}
