package dto

import (
	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/laundry/internal/domain"
)

type User struct {
	Email string `json:"email"`
}

type UserFilter struct {
	Search   *string `query:"search"`
	Limit    int     `query:"limit"`
	Page     int     `query:"page"`
	OrderBy  *string `query:"order_by"`
	OrderDir *string `query:"order_dir"`
}

func (f *UserFilter) ToDomain() (*domain.UserFilter, error) {
	// Ensure a minimum page of 1
	page := max(f.Page, 1)
	offset := (page - 1) * f.Limit

	var order *core.Order[domain.UserOrderField]
	if f.OrderBy != nil {
		dir := core.ASC
		if f.OrderDir != nil {
			dir = core.OrderDirection(*f.OrderDir)
		}

		order = &core.Order[domain.UserOrderField]{
			Field: domain.UserOrderField(*f.OrderBy),
			Dir:   dir,
		}

		err := order.Validate(func(uof domain.UserOrderField) bool {
			return uof == domain.CreatedAt || uof == domain.UpdatedAt
		})
		if err != nil {
			return nil, err
		}
	}

	filter := domain.NewUserFilter(
		f.Search,
		&core.Pagination{
			Limit:  f.Limit,
			Offset: offset,
		},
		order,
	)

	return &filter, nil
}
