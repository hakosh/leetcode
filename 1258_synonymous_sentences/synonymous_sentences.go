package synonymous_sentences

import (
	"sort"
	"strings"
)

// ["happy","joy"],["sad","sorrow"],["joy","cheerful"]

func generateSentences(synonyms [][]string, text string) []string {
	dict := make(map[string]*[]string, 0)

	for _, pair := range synonyms {
		first, hasFirst := dict[pair[0]]
		second, hasSecond := dict[pair[1]]

		if !hasFirst && !hasSecond {
			entry := make([]string, 2)
			copy(entry, pair)

			dict[pair[0]] = &entry
			dict[pair[1]] = &entry
		} else if hasFirst && hasSecond && first != second {
			for _, word := range *second {
				*first = append(*first, word)
				dict[word] = first
			}
		} else if hasFirst && !hasSecond {
			*first = append(*first, pair[1])
			dict[pair[1]] = first
		} else if !hasFirst && hasSecond {
			*second = append(*second, pair[0])
			dict[pair[0]] = second
		}
	}

	words := strings.Split(text, " ")

	sentences := make([]string, 0, len(words))
	generate(&sentences, words, dict, []string{})

	sort.Strings(sentences)

	return sentences
}

func generate(sentences *[]string, words []string, synonyms map[string]*[]string, sentence []string) {
	if len(words) == 0 {
		*sentences = append(*sentences, strings.Join(sentence, " "))
		return
	}

	if v, ok := synonyms[words[0]]; !ok {
		sentence = append(sentence, words[0])
		generate(sentences, words[1:], synonyms, sentence)
	} else {
		for _, synonym := range *v {
			newSentence := make([]string, len(sentence))
			copy(newSentence, sentence)
			newSentence = append(newSentence, synonym)
			generate(sentences, words[1:], synonyms, newSentence)
		}
	}
}
