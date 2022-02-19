package construct

import (
	"strings"
)

func CanConstruct(sentence string, words []string, memo map[string]bool) bool {
	if canConstruct, found := memo[sentence]; found {
		return canConstruct
	}
	if sentence == "" {
		return true
	}

	for _, word := range words {
		if strings.HasPrefix(sentence, word) {
			suffix := sentence[len(word):]
			if CanConstruct(suffix, words, memo) {
				memo[sentence] = true
				return true
			}
		}
	}

	memo[sentence] = false
	return false
}

func CanConstructTab(sentence string, words []string) bool {
	table := make([]bool, len(sentence)+1)
	table[0] = true
	for i := 0; i <= len(sentence); i++ {
		if table[i] {
			for _, word := range words {
				// if the word matches the characters starting at position i
				position := i + len(word)
				if position < len(table) && sentence[i:position] == word {
					table[position] = true
				}
			}
		}
	}

	return table[len(sentence)]
}
