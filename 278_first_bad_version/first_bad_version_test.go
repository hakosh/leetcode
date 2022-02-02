package first_bad_version

import "testing"

type Test struct {
	versions []bool
	out      int
}

var tests = []Test{
	{[]bool{false, false, false, true, true, true}, 4},
	{[]bool{true, true, true, true, true, true}, 1},
	{[]bool{false, false, false, false, false, true}, 6},
	{[]bool{false, false, false, false, false, true, true, true, true, true, true}, 6},
	{[]bool{false, false, true}, 3},
	{[]bool{true}, 1},
}

func TestFirstBadVersion(t *testing.T) {
	for _, test := range tests {
		setVersions(test.versions)

		if res := firstBadVersion(len(test.versions)); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.versions, test.out, res)
		}
	}
}
