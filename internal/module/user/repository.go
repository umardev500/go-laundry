package user

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	"github.com/umardev500/go-laundry/ent/profile"
	"github.com/umardev500/go-laundry/ent/tenant"
	userEntity "github.com/umardev500/go-laundry/ent/user"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/user"
)

type repositoryImpl struct {
	client *db.Client
}

// FindByToken implements user.Repository.
func (r *repositoryImpl) FindByToken(ctx context.Context, token string) (*user.User, error) {
	conn := r.client.GetConn(ctx)

	userEnt, err := conn.User.
		Query().
		Where(userEntity.ResetTokenEQ(token)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	var domainUser user.User
	r.mapFromEnt(userEnt, &domainUser)

	return &domainUser, nil
}

// Update implements user.Repository.
func (r *repositoryImpl) Update(ctx context.Context, payload *user.UserUpdate, userID uuid.UUID, tenantID *uuid.UUID) (*user.User, error) {
	conn := r.client.GetConn(ctx)

	// Fetch the user first
	u, err := conn.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Tenant scoping check
	if tenantID != nil {
		if u.TenantID != nil && *tenantID != *u.TenantID {
			return nil, fmt.Errorf("permission denied: cannot update user outside your tenant")
		}
	}

	// Prevent update if soft-deleted
	if u.DeletedAt != nil {
		return nil, fmt.Errorf("cannot update a deleted user")
	}

	// Start update builder
	userEnt, err := conn.User.UpdateOne(u).
		SetNillableEmail(payload.Email).
		SetNillablePassword(payload.Password).
		SetNillableResetToken(payload.ResetToken).
		SetNillableResetExpiresAt(payload.ResetExpiresAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// Mpa to domain user
	var domainUser user.User
	r.mapFromEnt(userEnt, &domainUser)

	return &domainUser, nil
}

func NewRepositoryImpl(client *db.Client) user.Repository {
	return &repositoryImpl{
		client: client,
	}
}

func (r *repositoryImpl) Create(ctx context.Context, u *user.UserCreate) (*user.User, error) {
	conn := r.client.GetConn(ctx)

	userReturned, err := conn.User.
		Create().
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetNillableTenantID(u.TenantID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	var result user.User
	r.mapFromEnt(userReturned, &result)
	return &result, nil
}

func (r *repositoryImpl) CreateProfile(ctx context.Context, userID uuid.UUID, u *user.ProfileCreate) (*user.Profile, error) {
	conn := r.client.GetConn(ctx)

	profile, err := conn.Profile.
		Create().
		SetUserID(userID).
		SetName(u.Name).
		SetNillableAvatar(u.Avatar).
		SetNillablePhone(u.Phone).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	domainProfile := user.Profile{
		ID:      profile.ID,
		Name:    *profile.Name,
		Avatar:  profile.Avatar,
		Phone:   profile.Phone,
		Created: profile.CreatedAt,
		Updated: profile.UpdatedAt,
	}

	return &domainProfile, err
}

func (r *repositoryImpl) Delete(ctx context.Context, tenantID *uuid.UUID, userID uuid.UUID) error {
	conn := r.client.GetConn(ctx)

	q := conn.User.
		Update().
		Where(userEntity.IDEQ(userID)).
		Where(userEntity.IDNEQ(userID))

	if tenantID != nil {
		q = q.Where(userEntity.TenantIDEQ(*tenantID))
	}

	// Soft delete
	_, err := q.SetDeletedAt(time.Now()).Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to soft delete user: %w", err)
	}

	return nil
}

func (r *repositoryImpl) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	conn := r.client.GetConn(ctx)
	u, err := conn.User.
		Query().
		Where(userEntity.EmailEQ(email)).
		Where(userEntity.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	var domainUser user.User
	r.mapFromEnt(u, &domainUser)

	return &domainUser, nil
}

// List implements user.Repository.
func (r *repositoryImpl) List(ctx context.Context, filter user.UserFilter) ([]*user.User, error) {
	conn := r.client.GetConn(ctx)

	// Start building query
	q := conn.User.Query()

	// Tenant scoping
	if filter.TenantID != nil {
		q = q.Where(userEntity.HasTenantWith(tenant.IDEQ(*filter.TenantID)))
	}

	// Soft delete filter
	if !filter.IncludeDeleted {
		q = q.Where(userEntity.DeletedAtIsNil())
	}

	// Serach by email or name
	if filter.Query != "" {
		q = q.Where(
			userEntity.Or(
				userEntity.EmailContainsFold(filter.Query),
				userEntity.HasProfileWith(profile.NameContainsFold(filter.Query)),
			),
		)
	}

	// Ordering
	switch filter.OrderBy {
	case user.OrderByEmailAsc:
		q = q.Order(ent.Asc(userEntity.FieldEmail))
	case user.OrderByEmailDesc:
		q = q.Order(ent.Desc(userEntity.FieldEmail))
	case user.OrderByCreatedAtAsc:
		q = q.Order(ent.Asc(userEntity.FieldCreatedAt))
	case user.OrderByCreatedAtDesc:
		q = q.Order(ent.Desc(userEntity.FieldCreatedAt))
	default:
		// default ordering
		q = q.Order(ent.Desc(userEntity.FieldCreatedAt))
	}

	// Pagination
	q = q.Limit(filter.Limit).Offset(filter.Offset)

	usersEnt, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	domainUsers := r.mapFromEnts(usersEnt)

	return domainUsers, nil
}

// PurgeUser implements user.Repository.
func (r *repositoryImpl) PurgeUser(ctx context.Context, tenantID *uuid.UUID, userID uuid.UUID) error {
	conn := r.client.GetConn(ctx)

	q := conn.User.
		Delete().
		Where(userEntity.IDEQ(userID)).
		Where(userEntity.IDNEQ(userID))

	if tenantID != nil {
		q = q.Where(userEntity.TenantIDEQ(*tenantID))
	}

	if _, err := q.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *repositoryImpl) UpdateProfile(ctx context.Context, userID uuid.UUID, u *user.ProfileUpdate) (*user.Profile, error) {
	conn := r.client.GetConn(ctx)

	profileEntity, err := conn.Profile.Query().
		Where(profile.HasUserWith(userEntity.IDEQ(userID))).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	profile, err := conn.Profile.
		UpdateOneID(profileEntity.ID).
		SetNillableName(u.Name).
		SetNillableAvatar(u.Avatar).
		SetNillablePhone(u.Phone).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	domainProfile := user.Profile{
		ID:      profile.ID,
		Name:    *profile.Name,
		Avatar:  profile.Avatar,
		Phone:   profile.Phone,
		Created: profile.CreatedAt,
		Updated: profile.UpdatedAt,
	}

	return &domainProfile, nil
}

func (r *repositoryImpl) mapFromEnt(e *ent.User, to *user.User) {
	if to == nil {
		return
	}

	to.ID = e.ID
	to.TenantID = e.TenantID
	to.Email = e.Email
	to.Password = e.Password
	to.ResetToken = e.ResetToken
	to.ResetExpiresAt = e.ResetExpiresAt
	to.CreatedAt = e.CreatedAt
	to.UpdatedAt = e.UpdatedAt
}

func (r *repositoryImpl) mapFromEnts(es []*ent.User) []*user.User {
	domainUsers := make([]*user.User, len(es))

	for i, e := range es {
		u := &user.User{}
		r.mapFromEnt(e, u)

		domainUsers[i] = u
	}

	return domainUsers
}
