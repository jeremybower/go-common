package pagination

import (
	"math"

	"github.com/jeremybower/go-common/guard"
)

type Result[T any] struct {
	PageIndex      int64
	PageSize       int64
	FirstItemIndex int64
	TotalItems     int64
	TotalPages     int64
	Items          []T
}

type Normalized struct {
	PageIndex      int64
	PageSize       int64
	FirstItemIndex int64
	TotalItems     int64
	TotalPages     int64
}

func Normalize(
	totalItems int64,
	pageIndex int64,
	pageSize int64,
	minimumPageSize int64,
	defaultPageSize int64,
	maximumPageSize int64,
) *Normalized {
	// Guard against invalid inputs.
	guard.GTE(totalItems, 0)
	guard.GTE(minimumPageSize, 1)
	guard.GTE(defaultPageSize, minimumPageSize)
	guard.GTE(maximumPageSize, defaultPageSize)

	// Normalize the page size.
	if pageSize == 0 {
		pageSize = defaultPageSize
	} else if pageSize < minimumPageSize {
		pageSize = minimumPageSize
	} else if pageSize > maximumPageSize {
		pageSize = maximumPageSize
	}

	// Calculate the total pages.
	totalPages := max(int64(1), int64(math.Ceil(float64(totalItems)/float64(pageSize))))

	// Normalize the page index.
	pageIndex = max(int64(0), min(pageIndex, totalPages-1))

	// Calculate the first item index.
	firstItemIndex := pageIndex * pageSize

	// Success.
	return &Normalized{
		PageIndex:      pageIndex,
		PageSize:       pageSize,
		FirstItemIndex: firstItemIndex,
		TotalItems:     totalItems,
		TotalPages:     totalPages,
	}
}
