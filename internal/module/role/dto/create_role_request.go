package dto

import "github.com/umardev500/go-laundry/internal/domain/role"

type CreateRoleRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}

func (r *CreateRoleRequest) ToRoleCreate() *role.RoleCreate {
	return &role.RoleCreate{
		Name:        r.Name,
		Description: r.Description,
	}
}
