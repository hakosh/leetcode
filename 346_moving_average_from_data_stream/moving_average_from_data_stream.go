package moving_average_from_data_stream

type MovingAverage struct {
	items []int
	size  int
	sum   int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		items: make([]int, 0),
		size:  size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.items) < this.size {
		this.sum += val
		this.items = append(this.items, val)
	} else {
		this.sum = this.sum - this.items[0] + val
		this.items = append(this.items[1:], val)
	}

	return float64(this.sum) / float64(len(this.items))
}
