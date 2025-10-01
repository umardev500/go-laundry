package dto

import "github.com/umardev500/go-laundry/internal/domain/user"

type UpdateProfileRequest struct {
	Name   *string `json:"name"`
	Avatar *string `json:"avatar"`
	Phone  *string `json:"phone" validate:"omitempty,min=10,max=15"`
}

func (u UpdateProfileRequest) ToUserProfileUpdate() *user.ProfileUpdate {
	return &user.ProfileUpdate{
		Name:   u.Name,
		Avatar: u.Avatar,
		Phone:  u.Phone,
	}
}
