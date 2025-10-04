package registration

import "github.com/umardev500/go-laundry/internal/domain/user"

type CreateUser struct {
	Profile *user.ProfileCreate
	User    *user.UserCreate
}
