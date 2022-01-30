package minimum_cost_for_tickets

import "testing"

type Test struct {
	days []int
	cost []int
	out  int
}

var tests = []Test{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{2, 4, 25}, 6},
	{[]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}, 11},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}, 17},
	{[]int{1, 4, 6, 7, 8, 20, 31, 32, 43, 47, 48, 52, 54, 60, 70, 75, 76, 77, 78, 98}, []int{2, 6, 12}, 32},
}

func TestMincostTickets(t *testing.T) {
	for _, test := range tests {
		if res := mincostTickets(test.days, test.cost); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.days, test.cost, test.out, res)
		}
	}
}

func BenchmarkMincostTickets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := tests[i%len(tests)]
		mincostTickets(t.days, t.cost)
	}
}
