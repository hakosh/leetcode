package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

var visited map[int]*Node

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited = make(map[int]*Node)

	return cloneNode(node)
}

func cloneNode(node *Node) *Node {
	if clone, ok := visited[node.Val]; ok {
		return clone
	}

	clone := &Node{Val: node.Val}
	visited[node.Val] = clone

	for _, neighbor := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, cloneNode(neighbor))
	}

	return clone
}
