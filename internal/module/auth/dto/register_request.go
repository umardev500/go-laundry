package dto

type RegisterRequest struct {
	Tenant TenantInfo `json:"tenant" validate:"required,dive"`
	Admin  AdminInfo  `json:"admin" validate:"required,dive"`
}

type TenantInfo struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Email   string `json:"email" validate:"email"`
	Address string `json:"address" validate:"required"`
}

type AdminInfo struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
}
