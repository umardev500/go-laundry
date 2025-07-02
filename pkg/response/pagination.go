package response

import (
	"math"
)

type PaginatedData[T any] struct {
	Items      []T            `json:"items"`
	Pagination PaginationInfo `json:"pagination"`
}

func NewPaginatedData[T any](items []T, page, limit, total int) PaginatedData[T] {
	pages := 1
	if limit > 0 {
		pages = int(math.Ceil(float64(total) / float64(limit)))
	}

	offset := (page - 1) * limit
	if page <= 0 {
		page = 1
		offset = 0
	}

	return PaginatedData[T]{
		Items: items,
		Pagination: PaginationInfo{
			Page:   page,
			Limit:  limit,
			Offset: offset,
			Total:  total,
			Pages:  pages,
		},
	}
}
