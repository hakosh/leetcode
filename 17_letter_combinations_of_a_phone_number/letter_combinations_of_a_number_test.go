package letter_combinations_of_a_phone_number

import (
	"reflect"
	"testing"
)

type Test struct {
	digits string
	out    []string
}

var tests = []Test{
	{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
}

func TestLetterCombinations(t *testing.T) {
	for _, test := range tests {
		if res := letterCombinations(test.digits); !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v expected %v, got %v", test.digits, test.out, res)
		}
	}
}

func BenchmarkLetterCombinations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinations("7979")
	}
}
