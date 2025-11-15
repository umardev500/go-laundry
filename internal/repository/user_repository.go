package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/profile"
	"github.com/umardev500/laundry/ent/user"
	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/laundry/internal/db"
	"github.com/umardev500/laundry/internal/domain"
	"github.com/umardev500/laundry/internal/errors"
)

type UserRepository interface {
	Create(ctx *core.Context, u *domain.User) (*domain.User, error)
	Find(ctx *core.Context, f domain.UserFilter) ([]*domain.User, int, error)
	FindByEmail(ctx *core.Context, email string) (*domain.User, error)
	FindByID(ctx *core.Context, id uuid.UUID) (*domain.User, error)
	Update(ctx *core.Context, u *domain.User) (*domain.User, error)
	UpdateProfile(ctx *core.Context, userID uuid.UUID, p *domain.Profile) error
}

type userRepositoryImpl struct {
	client *db.Client
}

func NewUserRepository(client *db.Client) UserRepository {
	return &userRepositoryImpl{
		client: client,
	}
}

// Create implements UserRepository.
func (r *userRepositoryImpl) Create(ctx *core.Context, u *domain.User) (*domain.User, error) {
	var user *ent.User

	err := r.client.WithTransaction(ctx, func(txCtx context.Context) error {
		var err error
		conn := r.client.GetConn(txCtx)

		user, err = conn.User.Create().
			SetEmail(u.Email).
			SetPassword(u.Password).
			Save(txCtx)
		if err != nil {
			return err
		}

		profile, err := conn.Profile.Create().
			SetUserID(user.ID).
			SetName(u.Profile.Name).
			Save(txCtx)
		if err != nil {
			return err
		}

		user.Edges.Profile = profile

		return nil
	})
	if err != nil {
		return nil, err
	}

	return r.mapEntToDomain(user), nil
}

// Find implements UserRepository
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

// FindByEmail implements UserRepository.
func (r *userRepositoryImpl) FindByEmail(ctx *core.Context, email string) (*domain.User, error) {
	conn := r.client.GetConn(ctx)
	q := conn.User.
		Query().
		Where(user.EmailEQ(email)).
		WithProfile()

	user, err := q.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NewUserNotFound(email)
		}

		return nil, err
	}

	return r.mapEntToDomain(user), nil
}

// FindByID implements UserRepository.
func (r *userRepositoryImpl) FindByID(ctx *core.Context, id uuid.UUID) (*domain.User, error) {
	conn := r.client.GetConn(ctx)
	q := conn.User.
		Query().
		Where(user.IDEQ(id)).
		WithProfile()

	user, err := q.First(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapEntToDomain(user), nil
}

// Update implements UserRepository.
// It returns user object without any edges.
func (r *userRepositoryImpl) Update(ctx *core.Context, u *domain.User) (*domain.User, error) {
	conn := r.client.GetConn(ctx)

	userObj, err := conn.User.UpdateOneID(u.ID).
		SetEmail(u.Email).
		SetPassword(u.Password).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapEntToDomain(userObj), nil
}

// UpdateProfile implements UserRepository.
func (r *userRepositoryImpl) UpdateProfile(ctx *core.Context, userID uuid.UUID, p *domain.Profile) error {
	conn := r.client.GetConn(ctx)

	err := conn.Profile.Update().
		SetName(p.Name).
		Where(profile.UserIDEQ(userID)).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// --- Helpers ---
func (r *userRepositoryImpl) mapEntToDomain(user *ent.User) *domain.User {
	if user == nil {
		return nil
	}

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
