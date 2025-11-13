package core

import (
	"math"
	"net/http"

	"github.com/umardev500/routerx"
)

type Response struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

type PaginationData struct {
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
	TotalItems int  `json:"total_items"`
	TotalPages int  `json:"total_pages"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
}

type PaginatedResponse struct {
	Data       any            `json:"data,omitempty"`
	Pagination PaginationData `json:"pagination"`
}

func NewSuccessResponse(c *routerx.Ctx, data any, status ...int) error {
	code := http.StatusOK

	if len(status) > 0 {
		code = status[0]
	}

	return c.Status(code).JSON(&Response{
		Data: data,
	})
}

func NewPaginatedResponse(c *routerx.Ctx, data any, page, pageSize, totalItems int, status ...int) error {
	code := http.StatusOK
	if len(status) > 0 {
		code = status[0]
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))

	return c.Status(code).JSON(&PaginatedResponse{
		Data: data,
		Pagination: PaginationData{
			Page:       page,
			PageSize:   pageSize,
			TotalItems: totalItems,
			TotalPages: totalPages,
			HasNext:    page < totalPages,
			HasPrev:    page > 1,
		},
	})
}

func NewErrorResponse(c *routerx.Ctx, err error, status ...int) error {
	code := http.StatusInternalServerError

	if len(status) > 0 {
		code = status[0]
	}

	return c.Status(code).JSON(&Response{
		Error: err.Error(),
	})
}
