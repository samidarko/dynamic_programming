package sum

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCanSum(t *testing.T) {
	assert.Equal(t, CanSum(7, []int{2, 3}, map[int]bool{}), true)
	assert.Equal(t, CanSum(7, []int{5, 3, 4, 7}, map[int]bool{}), true)
	assert.Equal(t, CanSum(7, []int{2, 4}, map[int]bool{}), false)
	assert.Equal(t, CanSum(8, []int{2, 3, 5}, map[int]bool{}), true)
	assert.Equal(t, CanSum(300, []int{7, 14}, map[int]bool{}), false)
}

func TestCanSumTab(t *testing.T) {
	assert.Equal(t, CanSumTab(7, []int{2, 3}), true)
	assert.Equal(t, CanSumTab(7, []int{5, 3, 4, 7}), true)
	assert.Equal(t, CanSumTab(7, []int{2, 4}), false)
	assert.Equal(t, CanSumTab(8, []int{2, 3, 5}), true)
	assert.Equal(t, CanSumTab(300, []int{7, 14}), false)
}
