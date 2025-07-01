package service

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-laundry/internal/domain"
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

// Login implements domain.AuthService.
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u.ID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * time.Duration(exp)).Unix(),
	})

	// Signed the token with the secret key
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		Token: tokenString,
	}, nil
}
