package validate_stack_sequences

import "testing"

type Test struct {
	pushed []int
	popped []int
	want   bool
}

var tests = []Test{
	{[]int{1}, []int{1}, true},
	{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, false},
	{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, true},
}

func TestValidateStackSequences(t *testing.T) {
	for _, test := range tests {
		if res := validateStackSequences(test.pushed, test.popped); res != test.want {
			t.Errorf("for %v and %v wanted %v, got %v", test.pushed, test.popped, test.want, res)
		}
	}
}
