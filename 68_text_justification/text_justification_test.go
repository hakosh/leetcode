package text_justification

import (
	"reflect"
	"testing"
)

type Test struct {
	words    []string
	maxWidth int
	want     []string
}

var tests = []Test{
	{
		[]string{"This", "is", "an", "example", "of", "text", "justification."},
		16,
		[]string{
			"This    is    an",
			"example  of text",
			"justification.  "},
	},
	{
		[]string{"What", "must", "be", "acknowledgment", "shall", "be"},
		16,
		[]string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        "},
	},
	{
		[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
		20,
		[]string{
			"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  "},
	},
}

func TestFullJustify(t *testing.T) {
	for _, test := range tests {
		if res := fullJustify(test.words, test.maxWidth); !reflect.DeepEqual(res, test.want) {
			t.Errorf("For %v and %v wanted %v, got %v", test.words, test.maxWidth, test.want, res)
		}
	}
}
