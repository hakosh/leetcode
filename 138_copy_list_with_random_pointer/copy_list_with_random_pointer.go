package copy_list_with_random_pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var visited map[*Node]*Node

func cloneNode(node *Node) *Node {
	if node == nil {
		return nil
	}

	if v, ok := visited[node]; ok {
		return v
	} else {
		newNode := &Node{Val: node.Val}
		visited[node] = newNode
		return newNode
	}
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	visited = make(map[*Node]*Node)

	oldNode := head
	newNode := &Node{Val: head.Val}

	visited[oldNode] = newNode

	for oldNode != nil {
		newNode.Random = cloneNode(oldNode.Random)
		newNode.Next = cloneNode(oldNode.Next)

		oldNode = oldNode.Next
		newNode = newNode.Next
	}

	return visited[head]
}
