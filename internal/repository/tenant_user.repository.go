package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/profile"
	"github.com/umardev500/laundry/ent/tenantuser"
	"github.com/umardev500/laundry/ent/user"
	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/laundry/internal/db"
	"github.com/umardev500/laundry/internal/domain"
)

type TenantUserRepository interface {
	Create(ctx *core.Context, tu *domain.TenantUser) (*domain.TenantUser, error)
	Delete(ctx *core.Context, id uuid.UUID) error
	Find(ctx *core.Context, f *domain.TenantUserFilter) ([]*domain.TenantUser, int, error)
	FindByID(core *core.Context, id uuid.UUID) (*domain.TenantUser, error)
	Updaate(ctx *core.Context, tu *domain.TenantUser) (*domain.TenantUser, error)
}

type tenantUserRepositoryImpl struct {
	client *db.Client
}

// Create implements TenantUserRepository.
func (r *tenantUserRepositoryImpl) Create(ctx *core.Context, tu *domain.TenantUser) (*domain.TenantUser, error) {
	panic("unimplemented")
}

// Delete implements TenantUserRepository.
func (r *tenantUserRepositoryImpl) Delete(ctx *core.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// Find implements TenantUserRepository.
func (r *tenantUserRepositoryImpl) Find(ctx *core.Context, f *domain.TenantUserFilter) ([]*domain.TenantUser, int, error) {
	q := r.client.GetConn(ctx).TenantUser.Query()
	criteria := f.Criteria

	// Apply search filter
	if criteria.Search != nil {
		search := *criteria.Search
		q = q.Where(
			tenantuser.Or(
				tenantuser.HasUserWith(user.EmailContainsFold(search)),
				tenantuser.HasUserWith(
					user.HasProfileWith(profile.NameContainsFold(search)),
				),
			),
		)
	}

	// Count total before pagination
	totalCount, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Include user (preload the edge)
	if criteria.IncludeUser {
		q = q.WithUser(func(uq *ent.UserQuery) {
			if criteria.IncludeProfile {
				uq.WithProfile()
			}
		})
	}

	// Apply pagination
	q = q.Offset(f.Pagination.Offset).Limit(f.Pagination.Limit)

	// Apply ordering
	orderStr := string(f.Order.Field)
	if f.Order.Dir == core.DESC {
		q = q.Order(ent.Desc(orderStr))
	} else {
		q = q.Order(ent.Asc(orderStr))
	}

	results, err := q.All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return r.mapEntToDomainList(results), totalCount, nil
}

// FindByID implements TenantUserRepository.
func (r *tenantUserRepositoryImpl) FindByID(core *core.Context, id uuid.UUID) (*domain.TenantUser, error) {
	panic("unimplemented")
}

// Updaate implements TenantUserRepository.
func (r *tenantUserRepositoryImpl) Updaate(ctx *core.Context, tu *domain.TenantUser) (*domain.TenantUser, error) {
	panic("unimplemented")
}

func NewTenantUserRepository(c *db.Client) TenantUserRepository {
	return &tenantUserRepositoryImpl{
		client: c,
	}
}

// --- Helpers ---
func (r *tenantUserRepositoryImpl) mapEntToDomain(tu *ent.TenantUser) *domain.TenantUser {
	var userName, userEmail string

	if tu.Edges.User != nil {
		userEmail = tu.Edges.User.Email

		if tu.Edges.User.Edges.Profile != nil {
			userName = tu.Edges.User.Edges.Profile.Name
		}
	}

	userOfTenant := domain.UserOfTenant{
		Name:  userName,
		Email: userEmail,
	}

	return &domain.TenantUser{
		ID:        tu.ID,
		TenantID:  tu.TenantID,
		UserID:    tu.UserID,
		User:      userOfTenant,
		CreatedAt: tu.CreatedAt,
		UpdatedAt: tu.UpdatedAt,
	}
}

func (r *tenantUserRepositoryImpl) mapEntToDomainList(tus []*ent.TenantUser) []*domain.TenantUser {
	var result []*domain.TenantUser
	for _, tu := range tus {
		result = append(result, r.mapEntToDomain(tu))
	}

	return result
}
