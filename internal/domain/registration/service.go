package registration

import (
	"context"

	"github.com/umardev500/go-laundry/internal/domain/user"
)

type Service interface {
	RegisterUser(ctx context.Context, payload *CreateUser) (*user.User, error)
}
