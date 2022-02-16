package excel_sheet_column_title

import "testing"

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{1, "A"},
	{27, "AA"},
	{701, "ZY"},
	{3340, "DXL"},
	{52, "AZ"},
	{53, "BA"},
}

func TestConvertToTitle(t *testing.T) {
	for _, test := range tests {
		if res := convertToTitle(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
