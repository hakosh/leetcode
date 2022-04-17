package jump_game_ii

import "testing"

type Test struct {
	nums []int
	want int
}

var tests = []Test{
	{[]int{2, 3, 1, 1, 4}, 2},
	{[]int{2, 3, 0, 1, 4}, 2},
	{[]int{1, 2, 1, 1, 1}, 3},
}

func TestJump(t *testing.T) {
	for _, test := range tests {
		if res := jump(test.nums); res != test.want {
			t.Errorf("for %v wanted %v, got %v", test.nums, test.want, res)
		}
	}
}
