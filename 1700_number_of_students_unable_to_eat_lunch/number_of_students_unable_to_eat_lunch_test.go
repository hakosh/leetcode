package number_of_students_unable_to_eat_lunch

import "testing"

type Test struct {
	students   []int
	sandwiches []int
	out        int
}

var tests = []Test{
	{[]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}, 3},
}

func TestCountStudents(t *testing.T) {
	for _, test := range tests {
		if res := countStudents(test.students, test.sandwiches); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.students, test.sandwiches, test.out, res)
		}
	}
}
