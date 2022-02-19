package graphs

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestConnectedComponentsCount(t *testing.T) {
	graph := map[int][]int{
		1: {2},
		2: {1},
		3: {},
		4: {6},
		5: {6},
		6: {4, 5, 7, 8},
		7: {6},
		8: {6},
	}

	assert.Equal(t, ConnectedComponentsCount(graph), 3)
}

func TestLargestComponent(t *testing.T) {
	graph := map[int][]int{
		0: {1, 5, 8},
		1: {0},
		2: {3, 4},
		3: {2, 4},
		4: {2, 3},
		5: {0, 8},
		8: {0, 5},
	}

	assert.Equal(t, LargestComponent(graph), 4)
}
