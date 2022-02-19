package construct

import (
	"strings"
)

func CountConstruct(sentence string, words []string, memo map[string]int) int {
	if countConstruct, found := memo[sentence]; found {
		return countConstruct
	}
	if sentence == "" {
		return 1
	}

	count := 0
	for _, word := range words {
		if strings.HasPrefix(sentence, word) {
			count += CountConstruct(sentence[len(word):], words, memo)
		}
	}

	memo[sentence] = count
	return count
}

func CountConstructTab(sentence string, words []string) int {
	table := make([]int, len(sentence)+1)
	table[0] = 1
	for i := 0; i <= len(sentence); i++ {
		for _, word := range words {
			// if the word matches the characters starting at position i
			position := i + len(word)
			if position <= len(sentence) && sentence[i:position] == word {
				table[position] += table[i]
			}
		}
	}

	return table[len(sentence)]
}
