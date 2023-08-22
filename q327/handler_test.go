package q327

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCountRangeSum(t *testing.T) {
	count := countRangeSum([]int{-2, 5, -1}, -2, 2)
	assert.Equal(t, 3, count)
}
