package repository

import "github.com/umardev500/go-laundry/internal/domain"

type userRepository struct{}

func NewUserRepository() domain.UserRepository {
	return &userRepository{}
}
