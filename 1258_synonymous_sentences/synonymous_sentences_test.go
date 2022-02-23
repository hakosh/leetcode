package synonymous_sentences

import (
	"fmt"
	"reflect"
	"testing"
)

type Test struct {
	synonyms [][]string
	text     string
	out      []string
}

var tests = []Test{
	{
		[][]string{{"happy", "joy"}, {"sad", "sorrow"}, {"joy", "cheerful"}},
		"I am happy today but was sad yesterday",
		[]string{
			"I am cheerful today but was sad yesterday",
			"I am cheerful today but was sorrow yesterday",
			"I am happy today but was sad yesterday",
			"I am happy today but was sorrow yesterday",
			"I am joy today but was sad yesterday",
			"I am joy today but was sorrow yesterday",
		},
	},
	{
		[][]string{{"c", "d"}, {"a", "b"}, {"b", "c"}, {"d", "e"}},
		"a b",
		[]string{
			"a a",
			"a b",
			"a c",
			"a d",
			"a e",
			"b a",
			"b b",
			"b c",
			"b d",
			"b e",
			"c a",
			"c b",
			"c c",
			"c d",
			"c e",
			"d a",
			"d b",
			"d c",
			"d d",
			"d e",
			"e a",
			"e b",
			"e c",
			"e d",
			"e e",
		},
	},
}

func TestGenerateSentences(t *testing.T) {
	for _, test := range tests {
		res := generateSentences(test.synonyms, test.text)

		fmt.Println(res)

		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("For %v and %v", test.synonyms, test.text)

			fmt.Println("Expected:")
			for _, sentence := range test.out {
				fmt.Println(sentence)
			}

			fmt.Println("\nGot:")
			for _, sentence := range res {
				fmt.Println(sentence)
			}
		}
	}
}
