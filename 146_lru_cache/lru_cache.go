package lru_cache

type Node struct {
	Val  [2]int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	size  int
	items map[int]*Node
	head  *Node
	tail  *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		size:  capacity,
		items: make(map[int]*Node, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.items[key]

	if !ok {
		return -1
	} else if len(this.items) == 1 || node == this.tail {
		return node.Val[1]
	} else if node == this.head {
		this.head = this.head.Next

		node.Prev = this.tail
		node.Next = nil
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		node.Next = nil
		node.Prev = this.tail
	}

	this.tail.Next = node
	this.tail = node

	return node.Val[1]
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		node := this.items[key]
		node.Val[1] = value
	} else {
		if len(this.items) == this.size {
			this.Evict()
		}

		node := &Node{
			Val:  [2]int{key, value},
			Next: nil,
			Prev: this.tail,
		}

		if this.tail != nil {
			this.tail.Next = node
		}

		if this.head == nil {
			this.head = node
		}

		this.items[key] = node
		this.tail = node
	}
}

func (this *LRUCache) Evict() {
	head := this.head

	if head != nil {
		this.head = head.Next
		delete(this.items, head.Val[0])
	}

	if this.tail == head {
		this.tail = nil
	}
}
