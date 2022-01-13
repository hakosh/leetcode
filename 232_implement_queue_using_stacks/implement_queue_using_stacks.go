package implement_queue_using_stacks

type MyQueue struct {
	items []int
	rev   []int
}

func Constructor() MyQueue {
	return MyQueue{
		items: make([]int, 0),
		rev:   make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	rev := make([]int, 0, len(this.items))

	for i := len(this.items) - 1; i >= 0; i-- {
		rev = append(rev, this.items[i])
		this.items = this.items[:i]
	}

	rev = append(rev, x)

	for i := len(rev) - 1; i >= 0; i-- {
		this.items = append(this.items, rev[i])
	}
}

func (this *MyQueue) Pop() int {
	top := this.items[len(this.items)-1]
	this.items = this.items[:len(this.items)-1]

	return top
}

func (this *MyQueue) Peek() int {
	return this.items[len(this.items)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.items) == 0
}
