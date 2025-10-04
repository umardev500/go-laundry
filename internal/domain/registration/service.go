package registration

import (
	"github.com/umardev500/go-laundry/internal/domain/user"

	appContext "github.com/umardev500/go-laundry/internal/app/context"
)

type Service interface {
	RegisterUser(ctx *appContext.ScopedContext, payload *CreateUser) (*user.User, error)
}
