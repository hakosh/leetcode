package populating_next_right_pointers_in_each_node_ii

import "testing"

var tree = &Node{
	Val: 1,
	Left: &Node{
		Val:   2,
		Left:  &Node{Val: 4},
		Right: &Node{Val: 5},
	},
	Right: &Node{
		Val:   3,
		Right: &Node{Val: 7},
	},
}

func TestConnect(t *testing.T) {
	connect(tree)

	connections := [][2]*Node{
		{tree.Left, tree.Right},
		{tree.Left.Left, tree.Left.Right},
		{tree.Left.Right, tree.Right.Right},
	}

	for _, conn := range connections {
		if conn[0].Next != conn[1] {
			t.Errorf("%v should be connected to %v, but it's connected to %v", conn[0].Val, conn[1].Val, conn[0].Next)
		}
	}
}
