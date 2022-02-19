package construct

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	assert.Equal(t, CanConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}, map[string]bool{}), true)
	assert.Equal(t, CanConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, map[string]bool{}), true)
	assert.Equal(t, CanConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, map[string]bool{}), false)
	assert.Equal(t, CanConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, map[string]bool{}), true)
	assert.Equal(t, CanConstruct("eeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee"}, map[string]bool{}), false)
}

func TestCanConstructTab(t *testing.T) {
	assert.Equal(t, CanConstructTab("purple", []string{"purp", "p", "ur", "le", "purpl"}), true)
	assert.Equal(t, CanConstructTab("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}), true)
	assert.Equal(t, CanConstructTab("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}), false)
	assert.Equal(t, CanConstructTab("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}), true)
	assert.Equal(t, CanConstructTab("eeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee"}), false)
}
