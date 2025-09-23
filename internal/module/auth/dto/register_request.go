package dto

type RegisterRequest struct {
	Tenant  TenantInfo  `json:"tenant" validate:"required,dive"`
	Profile UserProfile `json:"profile" validate:"required,dive"`
	Creds   Creds       `json:"admin" validate:"required,dive"`
}

type TenantInfo struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Email   string `json:"email" validate:"email"`
	Address string `json:"address" validate:"required"`
}

type UserProfile struct {
	Name string `json:"name" validate:"required"`
}

type Creds struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
}
