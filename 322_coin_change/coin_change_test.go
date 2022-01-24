package coin_change

import "testing"

type Test struct {
	coins  []int
	amount int
	out    int
}

var tests = []Test{
	{[]int{1, 2, 5}, 11, 3},
	{[]int{1, 2, 5}, 8762, 1753},
	{[]int{186, 419, 83, 408}, 6249, 20},
	{[]int{1, 2}, 11, 6},
	{[]int{2}, 3, -1},
	{[]int{1}, 0, 0},
}

func TestCoinChange(t *testing.T) {
	for _, test := range tests {
		if res := coinChange(test.coins, test.amount); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.coins, test.amount, test.out, res)
		}
	}
}

func BenchmarkCoinChange(b *testing.B) {
	coins := []int{186, 419, 83, 408}
	amount := 6249

	for i := 0; i < b.N; i++ {
		coinChange(coins, amount)
	}
}
