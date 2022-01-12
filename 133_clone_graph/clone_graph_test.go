package clone_graph

import (
	"testing"
)

var node1 = &Node{
	Val:       1,
	Neighbors: []*Node{node2, node4},
}

var node2 = &Node{
	Val: 2,
}

var node3 = &Node{
	Val:       3,
	Neighbors: []*Node{node2, node4},
}

var node4 = &Node{
	Val: 4,
}

func TestCloneGraph(t *testing.T) {
	node2.Neighbors = []*Node{node1, node3}
	node4.Neighbors = []*Node{node1, node3}

	c1 := cloneGraph(node1)

	if c1.Val != node1.Val {
		t.Errorf("Original: %v, clone: %v", node1.Val, c1.Val)
	}

	if len(c1.Neighbors) != len(node1.Neighbors) {
		t.Errorf("Children are not cloned")
	}

	for i, n := range c1.Neighbors {
		if n.Val != node1.Neighbors[i].Val {
			t.Errorf("Original: %v, clone: %v", node1.Neighbors[i].Val, n.Val)
		}
	}
}
