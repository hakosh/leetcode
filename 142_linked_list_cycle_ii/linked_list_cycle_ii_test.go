package linked_list_cycle_ii

import "testing"

type Test struct {
	in  *ListNode
	out *ListNode
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

var tests = []Test{
	{in: noCycle, out: nil},
	{in: node1, out: node2},
	{in: nil, out: nil},
}

func TestDetectCycle(t *testing.T) {
	node4.Next = node2

	for _, test := range tests {
		res := detectCycle(test.in)

		if res != test.out {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.out, res)
		}
	}
}
