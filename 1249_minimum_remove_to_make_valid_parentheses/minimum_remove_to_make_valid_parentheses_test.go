package minimum_remove_to_make_valid_parentheses

import "testing"

type Test struct {
	in  string
	out string
}

var tests = []Test{
	{"lee(t(c)o)de)", "lee(t(c)o)de"},
	{"a)b(c)d", "ab(c)d"},
	{"))((", ""},
	{"majom", "majom"},
	{"", ""},
	{"((()))", "((()))"},
	{"(A)", "(A)"},
	{"(A))", "(A)"},
	{"(((A)", "(A)"},
	{")(Z)((", "(Z)"},
	{"TdoivDDzZc(73hIOrswZ(tTdoivDD(zZc7)))3hIOrswZ)t", "TdoivDDzZc(73hIOrswZ(tTdoivDD(zZc7)))3hIOrswZt"},
}

func TestMinRemoveToMakeValid(t *testing.T) {
	for _, test := range tests {
		if res := minRemoveToMakeValid(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}

func BenchmarkMinRemoveToMakeValid(b *testing.B) {
	test := tests[len(tests)-1]
	for i := 0; i < b.N; i++ {
		minRemoveToMakeValid(test.in)
	}
}
