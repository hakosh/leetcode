package fizz_buzz

import (
	"reflect"
	"testing"
)

type Test struct {
	in  int
	out []string
}

var tests = []Test{
	{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
	{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
}

func TestFizzBuzz(t *testing.T) {
	for _, test := range tests {
		if res := fizzBuzz(test.in); !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
