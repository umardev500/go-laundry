package dto

type RequestPasswordResetRequest struct {
	Email string `json:"email" validate:"email"`
}
