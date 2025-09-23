package dto

import "github.com/umardev500/go-laundry/internal/domain/user"

type UpdateProfileRequest struct {
	Name    *string `json:"name"`
	Avatar  *string `json:"avatar"`
	Phone   *string `json:"phone" validate:"omitempty,min=10,max=15"`
	Address *string `json:"address" validate:"omitempty,min=10,max=100"`
}

func (u UpdateProfileRequest) ToUserProfileUpdate() *user.UserProfileUpdate {
	return &user.UserProfileUpdate{
		Name:    u.Name,
		Avatar:  u.Avatar,
		Phone:   u.Phone,
		Address: u.Address,
	}
}
