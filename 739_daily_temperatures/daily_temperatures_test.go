package daily_temperatures

import (
	"reflect"
	"testing"
)

type Test struct {
	in  []int
	out []int
}

var tests = []Test{
	{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
	{[]int{30, 60, 90}, []int{1, 1, 0}},
	{[]int{10, 9, 8, 7}, []int{0, 0, 0, 0}},
	{[]int{30, 30, 30, 30}, []int{0, 0, 0, 0}},
}

func TestDailyTemperatures(t *testing.T) {
	for _, test := range tests {
		if res := dailyTemperatures(test.in); !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}

func BenchmarkDailyTemperatures(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dailyTemperatures(tests[0].in)
	}
}
