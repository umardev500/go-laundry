package utils

import (
	"math"

	"github.com/umardev500/go-laundry/internal/types"
	"github.com/umardev500/go-laundry/pkg/response"
)

// CalculateTotalPages returns the number of pages based on total items and page size
func CalculateTotalPages(totalItems, pageSize int) int {
	if pageSize <= 0 {
		return 0
	}
	return int(math.Ceil(float64(totalItems) / float64(pageSize)))
}

func Paginate[T any](data []*T, total, offset, limit int) *types.PageResult[T] {
	return &types.PageResult[T]{
		Data: data,
		Pagination: &response.Pagination{
			Page:       offset + 1,
			PageSize:   limit,
			TotalItems: total,
			TotalPages: CalculateTotalPages(total, limit),
			HasNext:    offset+1 < CalculateTotalPages(total, limit),
			HasPrev:    offset > 1,
		},
	}
}
