package auth

import (
	"context"
	"time"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(ctx context.Context, email, password string) (user *user.User, token, refreshToken string, err error)
}

type serviceImpl struct {
	userRepo user.Repository
	cfg      *config.Config
}

func NewServiceImpl(userRepo user.Repository, cfg *config.Config) *serviceImpl {
	return &serviceImpl{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (s *serviceImpl) Login(ctx context.Context, email, password string) (user *user.User, token, refreshToken string, err error) {
	user, err = s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return
	}

	token, err = s.generateJWT(user)
	if err != nil {
		return
	}

	refreshToken, err = s.generateRefreshToken(user)
	if err != nil {
		return
	}

	return
}

func (s *serviceImpl) generateJWT(u *user.User) (tokenStr string, err error) {
	token, err := jwt.NewBuilder().
		Issuer(s.cfg.JWT.Issuer).
		Subject(u.ID.String()).
		IssuedAt(time.Now()).
		Expiration(time.Now().Add(time.Second * time.Duration(s.cfg.JWT.ExpirySeconds))).
		Build()
	if err != nil {
		return
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.HS256(), []byte(s.cfg.JWT.Secret)))
	if err != nil {
		return
	}

	tokenStr = string(signed)

	return
}

func (s *serviceImpl) generateRefreshToken(u *user.User) (tokenStr string, err error) {
	refreshToken, err := jwt.NewBuilder().
		Issuer(s.cfg.JWT.Issuer).
		Subject(u.ID.String()).
		IssuedAt(time.Now()).
		Expiration(time.Now().Add(time.Second * time.Duration(s.cfg.JWT.RefreshTokenExpirySeconds))).
		Build()
	if err != nil {
		return
	}

	signed, err := jwt.Sign(refreshToken, jwt.WithKey(jwa.HS256(), []byte(s.cfg.JWT.Secret)))
	if err != nil {
		return
	}

	tokenStr = string(signed)

	return
}
