package min_stack

type Node struct {
	val int
	min int
}

type MinStack struct {
	items []Node
}

func Constructor() MinStack {
	return MinStack{
		items: make([]Node, 0),
	}
}

func (this *MinStack) Push(val int) {
	node := Node{val: val, min: val}

	if len(this.items) > 0 {
		last := this.items[len(this.items)-1]

		if last.min < val {
			node.min = last.min
		}
	}

	this.items = append(this.items, node)
}

func (this *MinStack) Pop() {
	this.items = this.items[:len(this.items)-1]
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1].val
}

func (this *MinStack) GetMin() int {
	return this.items[len(this.items)-1].min
}
