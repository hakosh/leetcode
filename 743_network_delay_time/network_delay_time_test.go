package network_delay_time

import "testing"

type Test struct {
	times [][]int
	n     int
	k     int
	out   int
}

var tests = []Test{
	{[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
	{[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
	{[][]int{{1, 2, 1}, {3, 4, 1}}, 4, 1, -1},
	{[][]int{{1, 2, 9}, {1, 4, 2}, {2, 5, 1}, {4, 2, 4}, {4, 5, 6}, {3, 2, 3}, {5, 3, 7}, {3, 1, 5}}, 5, 1, 14},
}

func TestNetworkDelayTime(t *testing.T) {
	for _, test := range tests {
		if res := networkDelayTime(test.times, test.n, test.k); res != test.out {
			t.Errorf("For %v, %v and %v expected %v, got %v", test.times, test.n, test.k, test.out, res)
		}
	}
}

func BenchmarkNetworkDelayTime(b *testing.B) {
	test := tests[len(tests)-1]

	for i := 0; i < b.N; i++ {
		networkDelayTime(test.times, test.n, test.k)
	}
}
