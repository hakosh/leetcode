package shuffle_an_array

import "math/rand"

type Solution struct {
	nums     []int
	shuffled []int
}

func Constructor(nums []int) Solution {
	solution := Solution{nums: nums, shuffled: make([]int, len(nums))}
	copy(solution.shuffled, nums)

	return solution
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	copy(this.shuffled, this.nums)

	for i := 0; i < len(this.shuffled); i++ {
		idx := rand.Intn(len(this.shuffled))
		this.shuffled[i], this.shuffled[idx] = this.shuffled[idx], this.shuffled[i]
	}

	return this.shuffled
}
