package sum

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestBestSum(t *testing.T) {
	assert.Equal(t, BestSum(7, []int{5, 3, 4, 7}, map[int][]int{}), []int{7})
	assert.Equal(t, BestSum(8, []int{2, 3, 5}, map[int][]int{}), []int{5, 3})
	assert.Equal(t, BestSum(8, []int{1, 4, 5}, map[int][]int{}), []int{4, 4})
	assert.Equal(t, BestSum(100, []int{1, 2, 5, 25}, map[int][]int{}), []int{25, 25, 25, 25})
}

func TestBestSumTab(t *testing.T) {
	assert.Equal(t, BestSumTab(7, []int{5, 3, 4, 7}), []int{7})
	assert.Equal(t, BestSumTab(8, []int{2, 3, 5}), []int{3, 5})
	assert.Equal(t, BestSumTab(8, []int{1, 4, 5}), []int{4, 4})
	assert.Equal(t, BestSumTab(100, []int{1, 2, 5, 25}), []int{25, 25, 25, 25})
}
