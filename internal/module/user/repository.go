package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	"github.com/umardev500/go-laundry/ent/profile"
	userEntity "github.com/umardev500/go-laundry/ent/user"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/user"
)

type repositoryImpl struct {
	client *db.Client
}

func NewRepositoryImpl(client *db.Client) user.Repository {
	return &repositoryImpl{
		client: client,
	}
}

func (r *repositoryImpl) CreateUser(ctx context.Context, u *user.UserCreate) (*user.User, error) {
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

func (r *repositoryImpl) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	conn := r.client.GetConn(ctx)
	u, err := conn.User.
		Query().
		Where(userEntity.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	var domainUser user.User
	r.mapFromEnt(u, &domainUser)

	return &domainUser, nil
}

func (r *repositoryImpl) UpdateUserProfile(ctx context.Context, userID uuid.UUID, u *user.ProfileUpdate) (*user.Profile, error) {
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
		SetNillableAddress(u.Address).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	domainProfile := user.Profile{
		ID:      profile.ID,
		Name:    *profile.Name,
		Avatar:  profile.Avatar,
		Phone:   profile.Phone,
		Address: profile.Address,
		Created: profile.CreatedAt,
		Updated: profile.UpdatedAt,
	}

	return &domainProfile, nil
}

func (r *repositoryImpl) CreateUserProfile(ctx context.Context, userID uuid.UUID, u *user.ProfileCreate) (*user.Profile, error) {
	conn := r.client.GetConn(ctx)

	profile, err := conn.Profile.
		Create().
		SetUserID(userID).
		SetName(u.Name).
		SetNillableAvatar(u.Avatar).
		SetNillablePhone(u.Phone).
		SetNillableAddress(u.Address).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	domainProfile := user.Profile{
		ID:      profile.ID,
		Name:    *profile.Name,
		Avatar:  profile.Avatar,
		Phone:   profile.Phone,
		Address: profile.Address,
		Created: profile.CreatedAt,
		Updated: profile.UpdatedAt,
	}

	return &domainProfile, err
}

func (r *repositoryImpl) mapFromEnt(e *ent.User, to *user.User) {
	if to == nil {
		return
	}

	to.ID = e.ID
	to.Email = e.Email
	to.Password = e.Password
	to.ResetToken = e.ResetToken
	to.ResetExpiresAt = e.ResetExpiresAt
	to.CreatedAt = e.CreatedAt
	to.UpdatedAt = e.UpdatedAt
}
