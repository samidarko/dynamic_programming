package construct

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAllConstruct(t *testing.T) {
	assert.Equal(t, AllConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}, map[string][][]string{}), [][]string{{"purp", "le"}, {"p", "ur", "p", "le"}})
	assert.Equal(t, AllConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, map[string][][]string{}), [][]string{{"abc", "def"}})
	assert.Equal(t, AllConstruct("skateboard", []string{"ab", "cd", "abcd"}, map[string][][]string{}), [][]string{})
	assert.Equal(t, AllConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, map[string][][]string{}), [][]string{
		{"enter", "a", "p", "ot", "ent", "p", "ot"},
		{"enter", "a", "p", "ot", "ent", "p", "o", "t"},
		{"enter", "a", "p", "o", "t", "ent", "p", "ot"},
		{"enter", "a", "p", "o", "t", "ent", "p", "o", "t"},
	})
	assert.Equal(t, AllConstruct("eeeeeeeeeeeeeef", []string{"abc", "def"}, map[string][][]string{}), [][]string{})
}

func TestAllConstructTab(t *testing.T) {
	//var nilSlice [][]string
	//assert.Equal(t, AllConstructTab("purple", []string{"purp", "p", "ur", "le", "purpl"}), [][]string{{"purp", "le"}, {"p", "ur", "p", "le"}})
	//assert.Equal(t, AllConstructTab("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}), [][]string{{"abc", "def"}})
	//assert.Equal(t, AllConstructTab("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}), nilSlice)
	assert.Equal(t, AllConstructTab("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}), [][]string{
		{"enter", "a", "p", "ot", "ent", "p", "ot"},
		{"enter", "a", "p", "ot", "ent", "p", "o", "t"},
		{"enter", "a", "p", "o", "t", "ent", "p", "ot"},
		{"enter", "a", "p", "o", "t", "ent", "p", "o", "t"},
	})
	//assert.Equal(t, AllConstructTab("eeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee"}), nilSlice)
}
