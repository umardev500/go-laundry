package user

import (
	"context"

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
