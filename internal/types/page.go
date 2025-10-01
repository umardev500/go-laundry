package types

import "github.com/umardev500/go-laundry/pkg/response"

type PageData[T any] struct {
	Data  []*T
	Total int
}

type PageResult[T any] struct {
	Data       []*T
	Pagination *response.Pagination `json:"pagination,omitempty"`
}
