package graphs

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestIslandCount(t *testing.T) {
	grid := [][]string{
		{"W", "L", "W", "W", "W"},
		{"W", "L", "W", "W", "W"},
		{"W", "W", "W", "L", "W"},
		{"W", "W", "L", "L", "W"},
		{"L", "W", "W", "L", "L"},
		{"L", "L", "W", "W", "W"},
	}
	assert.Equal(t, IslandCount(grid), 3)
}
