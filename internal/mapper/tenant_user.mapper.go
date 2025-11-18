package mapper

import (
	"github.com/umardev500/laundry/internal/domain"
	"github.com/umardev500/laundry/internal/dto"
)

// MapDomainTenantUserToDTO maps a domain.TenantUser to dto.TenantUser
func MapDomainTenantUserToDTO(tu *domain.TenantUser) *dto.TenantUser {
	return &dto.TenantUser{
		ID:    tu.ID,
		Name:  tu.User.Name,
		Email: tu.User.Email,
	}
}

// MapDomainTenantUserToDTOs maps a slice of domain.TenantUser to a slice of dto.TenantUser
func MapDomainTenantUserToDTOs(tus []*domain.TenantUser) []*dto.TenantUser {
	tenants := make([]*dto.TenantUser, len(tus))
	for i, t := range tus {
		tenants[i] = MapDomainTenantUserToDTO(t)
	}
	return tenants
}
