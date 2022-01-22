package sum

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestHowSum(t *testing.T) {
	var nilSlice []int // temporary fix to assert slices are nil
	assert.Equal(t, HowSum(7, []int{2, 3}, map[int][]int{}), []int{3, 2, 2})
	assert.Equal(t, HowSum(7, []int{5, 3, 4, 7}, map[int][]int{}), []int{4, 3})
	assert.Equal(t, HowSum(7, []int{2, 4}, map[int][]int{}), nilSlice)
	assert.Equal(t, HowSum(8, []int{2, 3, 5}, map[int][]int{}), []int{2, 2, 2, 2})
	assert.Equal(t, HowSum(300, []int{7, 14}, map[int][]int{}), nilSlice)
}

func TestHowSumTab(t *testing.T) {
	var nilSlice []int // temporary fix to assert slices are nil
	assert.Equal(t, HowSumTab(7, []int{2, 3}), []int{3, 2, 2})
	assert.Equal(t, HowSumTab(7, []int{5, 3, 4, 7}), []int{4, 3})
	assert.Equal(t, HowSumTab(7, []int{2, 4}), nilSlice)
	assert.Equal(t, HowSumTab(8, []int{2, 3, 5}), []int{2, 2, 2, 2})
	assert.Equal(t, HowSumTab(300, []int{7, 14}), nilSlice)
}
