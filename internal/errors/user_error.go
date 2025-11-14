package errors

import (
	"fmt"
	"net/http"

	"github.com/umardev500/laundry/internal/core"
)

var (
	ErrUserNotFound      = fmt.Errorf("user not found")
	ErrUserAlreadyExists = fmt.Errorf("user already exists")
	ErrProfileRequired   = fmt.Errorf("user profile is required")
)

func NewUserNotFound(email string) *core.Error {
	return core.NewError(
		ErrUserNotFound,
		fmt.Sprintf("User with email %s not found", email),
		http.StatusNotFound,
	)
}

func NewUserAlreadyExists(email string) *core.Error {
	return core.NewError(
		ErrUserAlreadyExists,
		fmt.Sprintf("User with email %s already exists", email),
		http.StatusConflict,
	)
}

func NewErrProfileRequired() *core.Error {
	return core.NewError(
		ErrProfileRequired,
		"User profile is required",
		http.StatusBadRequest,
	)
}
