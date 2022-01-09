package design_circular_queue

type MyCircularQueue struct {
	items []int
	cap   int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		items: make([]int, 0, k),
		cap:   k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.items) == this.cap {
		return false
	}

	this.items = append(this.items, value)

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if len(this.items) == 0 {
		return false
	}

	this.items = this.items[1:]

	return true
}

func (this *MyCircularQueue) Front() int {
	if len(this.items) == 0 {
		return -1
	}

	return this.items[0]
}

func (this *MyCircularQueue) Rear() int {
	if len(this.items) == 0 {
		return -1
	}

	return this.items[len(this.items)-1]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.items) == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return len(this.items) == this.cap
}
