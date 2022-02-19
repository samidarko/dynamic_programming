package graphs

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestBuildGraph(t *testing.T) {
	edges := [][]string{
		{"i", "j"}, {"k", "i"},
		{"m", "k"}, {"k", "l"},
		{"o", "n"},
	}

	g, err := BuildGraph(edges)
	assert.Equal(t, err, nil)
	assert.Equal(t, g, map[string][]string{
		"i": {"j", "k"},
		"j": {"i"},
		"k": {"i", "m", "l"},
		"l": {"k"},
		"m": {"k"},
		"n": {"o"},
		"o": {"n"},
	})

	edge := []string{"only one node"}
	g, err = BuildGraph([][]string{edge})
	assert.Equal(t, g, map[string][]string{})
	assert.Equal(t, err != nil, true)
}
