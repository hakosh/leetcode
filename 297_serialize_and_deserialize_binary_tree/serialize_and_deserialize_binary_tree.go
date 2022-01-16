package serialize_and_deserialize_binary_tree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	s := strings.Builder{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			s.WriteString("_,")
		} else {
			s.WriteString(strconv.Itoa(node.Val) + ",")
			queue = append(queue, node.Left, node.Right)
		}
	}

	return s.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	parts := strings.Split(data, ",")

	if data == "" {
		return nil
	}

	v, _ := strconv.Atoi(parts[0])
	root := &TreeNode{Val: v}
	queue := []*TreeNode{root}

	i := 0

	for len(queue) > 0 && len(parts) > i {
		node := queue[0]
		queue = queue[1:]

		i++
		if parts[i] != "_" {
			v, _ := strconv.Atoi(parts[i])
			node.Left = &TreeNode{Val: v}
			queue = append(queue, node.Left)
		}

		i++
		if parts[i] != "_" {
			v, _ := strconv.Atoi(parts[i])
			node.Right = &TreeNode{Val: v}
			queue = append(queue, node.Right)
		}
	}

	return root
}

// DFS

type CodecDFS struct {
}

func ConstructorDFS() CodecDFS {
	return CodecDFS{}
}

// Serializes a tree to a single string.
func (this *CodecDFS) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	s := strings.Builder{}

	var h func(*TreeNode)
	h = func(root *TreeNode) {
		if root == nil {
			s.WriteString("_,")
			return
		} else {
			s.WriteString(strconv.Itoa(root.Val) + ",")
		}

		h(root.Left)
		h(root.Right)
	}

	h(root)

	return s.String()
}

// Deserializes your encoded data to tree.
func (this *CodecDFS) deserialize(data string) *TreeNode {
	parts := strings.Split(data, ",")

	if data == "" {
		return nil
	}

	i := 0

	var h func() *TreeNode
	h = func() *TreeNode {
		if parts[i] == "_" {
			i++
			return nil
		}

		v, _ := strconv.Atoi(parts[i])
		i++

		return &TreeNode{
			Val:   v,
			Left:  h(),
			Right: h(),
		}
	}

	return h()
}
