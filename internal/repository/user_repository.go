package repository

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/user"
	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/laundry/internal/db"
	"github.com/umardev500/laundry/internal/domain"
)

type UserRepository interface {
	Find(ctx *core.Context, f domain.UserFilter) ([]*domain.User, int, error)
}

type userRepositoryImpl struct {
	client *db.Client
}

func NewUserRepository(client *db.Client) *userRepositoryImpl {
	return &userRepositoryImpl{
		client: client,
	}
}

func (r *userRepositoryImpl) Find(ctx *core.Context, f domain.UserFilter) ([]*domain.User, int, error) {
	q := r.client.GetConn(ctx).User.Query()
	criteria := f.Filter

	// Apply search filter
	if criteria.Search != nil && *criteria.Search != "" {
		q = q.Where(
			user.Or(
				user.EmailContainsFold(*criteria.Search),
			),
		)
	}

	// Count total before pagination
	totalCount, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Include profile (preload the edge)
	if criteria.IncludeProfile {
		q = q.WithProfile()
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

// --- Helpers ---
func (r *userRepositoryImpl) mapEntToDomain(user *ent.User) *domain.User {
	var profile *domain.Profile
	if user.Edges.Profile != nil {
		profile = &domain.Profile{
			Name:      user.Edges.Profile.Name,
			CreatedAt: user.Edges.Profile.CreatedAt,
			UpdatedAt: user.Edges.Profile.UpdatedAt,
		}
	}

	return &domain.User{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		Profile:   profile,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (r *userRepositoryImpl) mapEntToDomainList(users []*ent.User) []*domain.User {
	var result []*domain.User
	for _, user := range users {
		result = append(result, r.mapEntToDomain(user))
	}
	return result
}
