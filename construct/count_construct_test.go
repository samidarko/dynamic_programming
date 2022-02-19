package construct

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCountConstruct(t *testing.T) {
	assert.Equal(t, CountConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}, map[string]int{}), 2)
	assert.Equal(t, CountConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, map[string]int{}), 1)
	assert.Equal(t, CountConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, map[string]int{}), 0)
	assert.Equal(t, CountConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, map[string]int{}), 4)
	assert.Equal(t, CountConstruct("eeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee"}, map[string]int{}), 0)
}

func TestCountConstructTab(t *testing.T) {
	assert.Equal(t, CountConstructTab("purple", []string{"purp", "p", "ur", "le", "purpl"}), 2)
	assert.Equal(t, CountConstructTab("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}), 1)
	assert.Equal(t, CountConstructTab("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}), 0)
	assert.Equal(t, CountConstructTab("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}), 4)
	assert.Equal(t, CountConstructTab("eeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee"}), 0)
}
