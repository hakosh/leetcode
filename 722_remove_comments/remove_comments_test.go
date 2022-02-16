package remove_comments

import (
	"reflect"
	"testing"
)

type Test struct {
	code []string
	out  []string
}

var tests = []Test{
	{
		[]string{"/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"},
		[]string{"int main()", "{ ", "  ", "int a, b, c;", "a = b + c;", "}"},
	},
	{
		[]string{"int a; /* cut this */ int b; /* and this */"},
		[]string{"int a;  int b; "},
	},
	{
		[]string{"a/*comment", "line", "more_comment*/b"},
		[]string{"ab"},
	},
}

func TestRemoveComments(t *testing.T) {
	for i, test := range tests {
		if res := removeComments(test.code); !reflect.DeepEqual(res, test.out) {
			t.Errorf("For test %d got %v", i, res)
		}
	}
}
