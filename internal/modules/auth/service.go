package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/umardev500/go-laundry/internal/config"
)

type AuthService interface {
	CreateToken(ctx context.Context) (string, error)
	VerifyToken(token string) (*Claims, error)
}

type authService struct {
	config *config.AppConfig
}

func NewAuthService(appConfig *config.AppConfig) AuthService {
	return &authService{
		config: appConfig,
	}
}

func (a *authService) CreateToken(ctx context.Context) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    a.config.Jwt.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.config.Jwt.Expiration)),
		},
	})

	return token.SignedString([]byte(a.config.Jwt.Secret))
}

func (a *authService) VerifyToken(token string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(a.config.Jwt.Secret), nil
	})

	return claims, err
}
