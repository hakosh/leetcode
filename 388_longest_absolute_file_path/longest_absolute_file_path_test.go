package longest_absolute_file_path

import "testing"

type Test struct {
	input string
	out   int
}

var tests = []Test{
	{"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext", 32},
	{"a", 0},
	{"", 0},
	{"majom.txt", 9},
}

func TestLenghtLongestPath(t *testing.T) {
	for _, test := range tests {
		if res := lengthLongestPath(test.input); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.input, test.out, res)
		}
	}
}
