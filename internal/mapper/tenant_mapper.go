package mapper

import (
	"github.com/umardev500/laundry/internal/domain"
	"github.com/umardev500/laundry/internal/dto"
)

// MapDomainTenantToDTO maps a domain.Tenant to dto.Tenant
func MapDomainTenantToDTO(t *domain.Tenant) *dto.Tenant {
	return &dto.Tenant{
		ID:   t.ID,
		Name: t.Name,
	}
}

// MapDomainTenantToDTOs map a slice of  domain.Tenant to a slice of dto.Tenant.
func MapDomainTenantToDTOs(ts []*domain.Tenant) []*dto.Tenant {
	tenants := make([]*dto.Tenant, len(ts))
	for i, t := range ts {
		tenants[i] = MapDomainTenantToDTO(t)
	}
	return tenants
}
