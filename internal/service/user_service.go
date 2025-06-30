package service

import "github.com/umardev500/go-laundry/internal/domain"

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{
		repo: repo,
	}
}
