package design_linked_list

type Node struct {
	val  int
	prev *Node
	next *Node
}

type MyLinkedList struct {
	len  int
	head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.len {
		return -1
	}

	node := this.head

	for i := 0; i < index; i++ {
		node = node.next
	}

	return node.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{val: val, next: this.head}

	if this.head != nil {
		this.head.prev = node
	}

	this.head = node
	this.len++
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.AddAtHead(val)
		return
	}

	last := this.head

	for last.next != nil {
		last = last.next
	}

	node := &Node{val: val, prev: last}
	last.next = node
	this.len++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.len {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.len {
		this.AddAtTail(val)
		return
	}

	before := this.head
	for i := 1; i < index; i++ {
		before = before.next
	}
	after := before.next

	node := &Node{val: val, prev: before, next: after}

	after.prev = node
	before.next = node
	this.len++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.len {
		return
	}

	if index == 0 {
		this.head = this.head.next
		this.len--
		return
	}

	node := this.head
	for i := 0; i < index; i++ {
		node = node.next
	}

	if node.next == nil {
		node.prev.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	node.prev = nil
	node.next = nil

	this.len--
}
