package service

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-laundry/internal/domain"
	"github.com/umardev500/go-laundry/internal/ent"
	sharedjwt "github.com/umardev500/go-laundry/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo domain.UserRepository
}

func NewAuthService(userRepo domain.UserRepository) domain.AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (a *authService) Login(ctx context.Context, payload *domain.LoginRequest) (*domain.LoginResponse, error) {
	u, err := a.userRepo.GetByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}

	// Compare the input password with the stored hash
	inputPass := []byte(payload.Password)
	hashedPass := []byte(u.PasswordHash)

	err = bcrypt.CompareHashAndPassword(hashedPass, inputPass)
	if err != nil {
		return nil, err
	}

	// Get JWT configuration from env
	secret := os.Getenv("JWT_SECRET")
	expireAt := os.Getenv("JWT_EXPIRATION_HOURS")

	if secret == "" || expireAt == "" {
		log.Fatal().Err(err).Msg("JWT_SECRET or JWT_EXPIRATION_HOURS is not set")
	}

	exp, err := strconv.Atoi(expireAt)
	if err != nil {
		return nil, err
	}

	// Create a new JWT token with claims
	claims := domain.Claims{
		Sub:      u.ID,
		Merchant: u.Edges.Merchants.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(exp))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token, err := sharedjwt.Sign(claims, []byte(secret))
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		Token: token,
	}, nil
}

func (a *authService) Me(ctx context.Context) (*ent.User, error) {
	claims, err := sharedjwt.Claims[*domain.Claims](ctx)
	if err != nil {
		return nil, err
	}

	return a.userRepo.GetByID(ctx, claims.Sub)
}
