package graphs

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestShortestPath(t *testing.T) {
	edges := [][]string{
		{"w", "x"},
		{"x", "y"},
		{"z", "y"},
		{"z", "v"},
		{"w", "v"},
	}

	assert.Equal(t, ShortestPath(edges, "w", "z"), 2)
	assert.Equal(t, ShortestPath(edges, "w", "a"), -1)
}
