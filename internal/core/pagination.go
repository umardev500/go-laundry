package core

import "fmt"

type OrderDirection string

const (
	ASC  OrderDirection = "asc"
	DESC OrderDirection = "desc"
)

type Order[T any] struct {
	Field T
	Dir   OrderDirection
}

func (o *Order[T]) Validate(isValidField func(T) bool) error {
	// Validte direction
	switch o.Dir {
	case ASC, DESC:
	default:
		return fmt.Errorf("invalid order direction: %s", o.Dir)
	}

	// Validate field based on type T
	if !isValidField(o.Field) {
		return fmt.Errorf("invalid order field: %v", o.Field)
	}

	return nil
}

type Pagination struct {
	Limit  int
	Offset int
}

func DefaultOrderFallback[T any](order *Order[T], fallbackField T) Order[T] {
	if order != nil {
		return *order
	}

	return Order[T]{
		Field: fallbackField,
		Dir:   ASC,
	}
}

func DefaultPaginationFallback(paging *Pagination) Pagination {
	if paging != nil {
		return *paging
	}

	return Pagination{
		Limit:  10,
		Offset: 0,
	}
}
