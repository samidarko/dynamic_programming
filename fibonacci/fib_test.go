package fibonacci

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestFib(t *testing.T) {
	assert.Equal(t, Fib(6, Memo{}), 8)
	assert.Equal(t, Fib(7, Memo{}), 13)
	assert.Equal(t, Fib(8, Memo{}), 21)
	assert.Equal(t, Fib(50, Memo{}), 12586269025)
}

func TestFibTab(t *testing.T) {
	assert.Equal(t, FibTab(6), 8)
	assert.Equal(t, FibTab(7), 13)
	assert.Equal(t, FibTab(8), 21)
	assert.Equal(t, FibTab(50), 12586269025)
}
