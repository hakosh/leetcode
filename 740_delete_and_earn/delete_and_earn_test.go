package delete_and_earn

import "testing"

type Test struct {
	in  []int
	out int
}

var tests = []Test{
	{[]int{1}, 1},
	{[]int{2, 3, 4}, 6},
	{[]int{2, 2, 3, 3, 3, 4}, 9},
	{[]int{2, 5, 5, 6, 8}, 20},
	{[]int{8, 10, 4, 9, 1, 3, 5, 9, 4, 10}, 37},
	{[]int{1, 3, 4, 4, 5, 8, 9, 9, 10, 10}, 37},
	{[]int{12, 32, 93, 17, 100, 72, 40}, 366},
	{[]int{1, 1, 1, 2, 4, 5, 5, 5, 6}, 18},
}

func TestDeleteAndEarn(t *testing.T) {
	for _, test := range tests {
		if res := deleteAndEarn(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}

func BenchmarkDeleteAndEarn(b *testing.B) {
	nums := []int{12, 32, 93, 17, 100, 72, 40}

	for i := 0; i < b.N; i++ {
		deleteAndEarn(nums)
	}
}
