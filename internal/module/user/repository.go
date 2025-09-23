package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent/profile"
	userEntity "github.com/umardev500/go-laundry/ent/user"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/user"
)

type repositoryImpl struct {
	client *db.Client
}

func NewRepositoryImpl(client *db.Client) *repositoryImpl {
	return &repositoryImpl{
		client: client,
	}
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
	domainUser.MapFromEnt(u)

	return &domainUser, nil
}

func (r *repositoryImpl) UpdateUserProfile(ctx context.Context, userID uuid.UUID, u *user.UserProfileUpdate) (*user.UserProfile, error) {
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

	domainProfile := user.UserProfile{
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

func (r *repositoryImpl) CreateUserProfile(ctx context.Context, u *user.UserProfileCreate) (*user.UserProfile, error) {
	conn := r.client.GetConn(ctx)

	profile, err := conn.Profile.
		Create().
		SetName(u.Name).
		SetNillableAvatar(u.Avatar).
		SetNillablePhone(u.Phone).
		SetNillableAddress(u.Address).
		Save(ctx)

	domainProfile := user.UserProfile{
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
