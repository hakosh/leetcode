package design_linked_list

type Node struct {
	val  int
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
	this.head = node
	this.len++
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.AddAtHead(val)
		return
	}

	last := this.head

	for {
		if last.next != nil {
			last = last.next
		} else {
			break
		}
	}

	node := &Node{val: val}
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

	before := this.head
	for i := 1; i < index; i++ {
		before = before.next
	}

	node := &Node{val: val, next: before.next}
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

	before := this.head
	for i := 1; i < index; i++ {
		before = before.next
	}

	before.next = before.next.next
	this.len--
}
