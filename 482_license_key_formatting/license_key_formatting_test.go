package license_key_formatting

import "testing"

type Test struct {
	s   string
	k   int
	out string
}

var tests = []Test{
	{"5F3Z-2e-9-w", 4, "5F3Z-2E9W"},
	{"2-5g-3-J", 2, "2-5G-3J"},
	{"---", 3, ""},
}

func TestLicenseKeyFormatting(t *testing.T) {
	for _, test := range tests {
		if res := licenseKeyFormatting(test.s, test.k); res != test.out {
			t.Errorf("For %v and %v expected %q, got %q", test.s, test.k, test.out, res)
		}
	}
}
