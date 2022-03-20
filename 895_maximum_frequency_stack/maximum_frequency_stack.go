package maximum_frequency_stack

type FreqStack struct {
	freq    map[int]int
	group   map[int][]int
	maxFreq int
}

func Constructor() FreqStack {
	return FreqStack{
		freq:  make(map[int]int),
		group: make(map[int][]int),
	}
}

func (this *FreqStack) Push(val int) {
	freq := 1
	if f, ok := this.freq[val]; ok {
		freq = f + 1
	}

	this.freq[val] = freq

	if freq > this.maxFreq {
		this.maxFreq = freq
	}

	if g, ok := this.group[freq]; ok {
		this.group[freq] = append(g, val)
	} else {
		this.group[freq] = []int{val}
	}
}

func (this *FreqStack) Pop() int {
	g := this.group[this.maxFreq]
	x := g[len(g)-1]
	g = g[:len(g)-1]
	this.group[this.maxFreq] = g
	this.freq[x] = this.freq[x] - 1

	if len(g) == 0 {
		this.maxFreq--
	}

	return x
}
