package domain

type CreateFeatureInput struct {
	Name        string
	Description string
	Enabled     bool
	Permissions []CreatePermissionInput
}
