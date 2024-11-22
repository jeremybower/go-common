package pagination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalize(t *testing.T) {
	normalized := Normalize(100, 0, 10, 2, 10, 100)
	assert.Equal(t, int64(0), normalized.PageIndex)
	assert.Equal(t, int64(10), normalized.PageSize)
	assert.Equal(t, int64(0), normalized.FirstItemIndex)
	assert.Equal(t, int64(100), normalized.TotalItems)
	assert.Equal(t, int64(10), normalized.TotalPages)
}

func TestNormalizeWhenTotalItemsIsNegative(t *testing.T) {
	assert.Panics(t, func() {
		Normalize(-1, 0, 10, 2, 10, 100)
	})
}

func TestNormalizeWhenMinimumPageSizeIsLessThanOne(t *testing.T) {
	assert.Panics(t, func() {
		Normalize(100, 0, 10, 0, 10, 100)
	})
}

func TestNormalizeWhenDefaultPageSizeIsLessThanMinimumPageSize(t *testing.T) {
	assert.Panics(t, func() {
		Normalize(100, 0, 10, 2, 1, 100)
	})
}

func TestNormalizeWhenMaximumPageSizeIsLessThanDefaultPageSize(t *testing.T) {
	assert.Panics(t, func() {
		Normalize(100, 0, 10, 2, 10, 9)
	})
}

func TestNormalizeWhenPageSizeIsZero(t *testing.T) {
	normalized := Normalize(100, 0, 0, 2, 10, 100)
	assert.Equal(t, int64(10), normalized.PageSize)
}

func TestNormalizeWhenPageSizeIsLessThanMinimumPageSize(t *testing.T) {
	normalized := Normalize(100, 0, 1, 2, 10, 100)
	assert.Equal(t, int64(2), normalized.PageSize)
}

func TestNormalizeWhenPageSizeIsGreaterThanMaximumPageSize(t *testing.T) {
	normalized := Normalize(100, 0, 101, 2, 10, 100)
	assert.Equal(t, int64(100), normalized.PageSize)
}
