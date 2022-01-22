package traveler

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGridTraveler(t *testing.T) {
	assert.Equal(t, GridTraveler(1, 1, map[string]int{}), 1)
	assert.Equal(t, GridTraveler(2, 3, map[string]int{}), 3)
	assert.Equal(t, GridTraveler(3, 2, map[string]int{}), 3)
	assert.Equal(t, GridTraveler(3, 3, map[string]int{}), 6)
	assert.Equal(t, GridTraveler(18, 18, map[string]int{}), 2333606220)
}

func TestGridTravelerTab(t *testing.T) {
	assert.Equal(t, GridTravelerTab(1, 1), 1)
	assert.Equal(t, GridTravelerTab(2, 3), 3)
	assert.Equal(t, GridTravelerTab(3, 2), 3)
	assert.Equal(t, GridTravelerTab(3, 3), 6)
	assert.Equal(t, GridTravelerTab(18, 18), 2333606220)
}
