package boats_to_save_people

import "testing"

type Test struct {
	people []int
	limit  int
	want   int
}

var tests = []Test{
	{[]int{1, 2}, 3, 1},
	{[]int{3, 2, 2, 1}, 3, 3},
	{[]int{3, 5, 3, 4}, 5, 4},
	{[]int{3, 2, 3, 2, 2}, 6, 3},
}

func TestNumRescueBoats(t *testing.T) {
	for _, test := range tests {
		if res := numRescueBoats(test.people, test.limit); res != test.want {
			t.Errorf("for %v and %v wanted %v, got %v", test.people, test.limit, test.want, res)
		}
	}
}
