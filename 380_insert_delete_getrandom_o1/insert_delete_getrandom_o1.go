package insert_delete_getrandom_o1

import (
	"math/rand"
)

type RandomizedSet struct {
	items []int
	dict  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		items: make([]int, 0),
		dict:  make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dict[val]; ok {
		return false
	}

	this.dict[val] = len(this.items)
	this.items = append(this.items, val)

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.dict[val]

	if !ok {
		return false
	}

	lastIdx := len(this.items) - 1
	lastVal := this.items[lastIdx]

	// Switch
	this.items[idx] = lastVal
	this.dict[lastVal] = idx

	// Remove
	this.items = this.items[:lastIdx]
	delete(this.dict, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.items))
	return this.items[idx]
}
