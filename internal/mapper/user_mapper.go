package mapper

import (
	"github.com/umardev500/laundry/internal/domain"
	"github.com/umardev500/laundry/internal/dto"
)

// MapDomainUserToDTO maps a domain.User to dto.User
func MapDomainUserToDTO(u *domain.User) *dto.User {
	if u == nil {
		return nil
	}
	return &dto.User{
		Email: u.Email,
	}
}

// MapDomainUsersToDTOs maps a slice of domain.User to a slice of dto.User
func MapDomainUsersToDTOs(users []*domain.User) []*dto.User {
	dtos := make([]*dto.User, len(users))
	for i, u := range users {
		dtos[i] = MapDomainUserToDTO(u)
	}
	return dtos
}
