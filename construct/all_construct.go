package construct

import (
	"strings"
)

func AllConstruct(sentence string, words []string, memo map[string][][]string) [][]string {
	if constructs, found := memo[sentence]; found {
		return constructs
	}
	if sentence == "" {
		return [][]string{{}}
	}

	result := make([][]string, 0)

	for _, word := range words {
		if strings.HasPrefix(sentence, word) {
			ways := AllConstruct(sentence[len(word):], words, memo)
			for i := range ways {
				ways[i] = append([]string{word}, ways[i]...)
				memo[sentence] = result
			}
			result = append(result, ways...)
		}
	}

	return result
}

func AllConstructTab(sentence string, words []string) [][]string {
	table := make([][][]string, len(sentence)+2)
	table[0] = [][]string{{}}
	for i := 0; i <= len(sentence); i++ {
		for _, word := range words {
			position := i + len(word)
			if position > len(sentence) {
				position = len(sentence)
			}
			if sentence[i:position] == word {
				ways := make([][]string, len(table[i]))
				for j := range table[i] {
					ways[j] = append(table[i][j], word)
				}
				table[position] = append(table[position], ways...)
			}
		}
	}

	return table[len(sentence)]
}
