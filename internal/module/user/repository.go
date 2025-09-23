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

func (r *repositoryImpl) UpsertUserProfile(ctx context.Context, userID uuid.UUID, input *user.UserProfileUpsert) (*user.UserProfile, error) {
	conn := r.client.GetConn(ctx)
	existing, err := conn.Profile.
		Query().
		Where(profile.HasUserWith(userEntity.IDEQ(userID))).
		Only(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if existing != nil {
		return r.updateUserProfile(ctx, existing, input.Update)
	}

	return r.createUserProfile(ctx, input.Create)
}

func (r *repositoryImpl) updateUserProfile(ctx context.Context, existing *ent.Profile, u *user.UserProfileUpdate) (*user.UserProfile, error) {
	profile, err := existing.Update().
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

func (r *repositoryImpl) createUserProfile(ctx context.Context, u *user.UserProfileCreate) (*user.UserProfile, error) {
	profile, err := r.client.GetConn(ctx).Profile.
		Create().
		SetName(u.Name).
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
