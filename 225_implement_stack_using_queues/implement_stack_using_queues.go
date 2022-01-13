package implement_stack_using_queues

type MyStack struct {
	items []int
}

func Constructor() MyStack {
	return MyStack{
		items: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.items = append(this.items, x)
}

func (this *MyStack) Pop() int {
	tmp := make([]int, 0, len(this.items))

	for len(this.items) > 1 {
		tmp = append(tmp, this.items[0])
		this.items = this.items[1:]
	}

	value := this.items[0]
	this.items = tmp

	return value
}

func (this *MyStack) Top() int {
	value := this.Pop()
	this.Push(value)
	return value
}

func (this *MyStack) Empty() bool {
	return len(this.items) == 0
}
