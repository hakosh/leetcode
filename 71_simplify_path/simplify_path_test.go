package simplify_path

import "testing"

type Test struct {
	path string
	want string
}

var tests = []Test{
	{"/", "/"},
	{"/_/_", "/_/_"},
	{"/../", "/"},
	{"/..", "/"},
	{"/home/", "/home"},
	{"/home//foo", "/home/foo"},
	{"/home///foo//", "/home/foo"},
	{"/path/../hello/./world", "/hello/world"},
}

func TestSimplifyPath(t *testing.T) {
	for _, test := range tests {
		if res := simplifyPath(test.path); res != test.want {
			t.Errorf("For %v wanted %v, got %v", test.path, test.want, res)
		}
	}
}
