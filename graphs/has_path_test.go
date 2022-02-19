package graphs

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestHasPathRecursive(t *testing.T) {
	graph := map[string][]string{
		"a": {"c", "b"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	assert.Equal(t, HasPathRecursive(graph, "a", "f"), true)
	assert.Equal(t, HasPathRecursive(graph, "e", "f"), false)
}

func TestHasPathIterative(t *testing.T) {
	graph := map[string][]string{
		"a": {"c", "b"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	assert.Equal(t, HasPathIterative(graph, "a", "f"), true)
	assert.Equal(t, HasPathIterative(graph, "e", "f"), false)
}

func TestHasPathUndirected(t *testing.T) {
	edges := [][]string{
		{"i", "j"}, {"k", "i"},
		{"m", "k"}, {"k", "l"},
		{"o", "n"},
	}

	g, _ := BuildGraph(edges)
	assert.Equal(t, HasPathUndirected(g, "i", "l", map[string]bool{}), true)
	assert.Equal(t, HasPathUndirected(g, "i", "o", map[string]bool{}), false)
}
