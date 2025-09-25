package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/domain/user"
	"github.com/umardev500/go-laundry/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(ctx context.Context, email, password string) (user *user.User, token, refreshToken string, err error)
	ResetPassword(ctx context.Context, token, newPassword string) (user *user.User, accessToken, refreshToken string, err error)
	RequestPasswordReset(ctx context.Context, email string) error
}

type serviceImpl struct {
	userService user.Service
	cfg         *config.Config
}

func NewServiceImpl(userService user.Service, cfg *config.Config) *serviceImpl {
	return &serviceImpl{
		userService: userService,
		cfg:         cfg,
	}
}

func (s *serviceImpl) Login(ctx context.Context, email, password string) (user *user.User, token, refreshToken string, err error) {
	user, err = s.userService.FindByEmail(ctx, email)
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

func (s *serviceImpl) ResetPassword(ctx context.Context, token, newPassword string) (u *user.User, accessToken, refreshToken string, err error) {
	// Validate token, get user
	userData, err := s.validateResetToken(ctx, token)
	if err != nil {
		return
	}

	// Prepare payload
	payload := &user.UserUpdate{
		Password: &newPassword,
	}

	// Call user service to update credentials
	updateduser, err := s.userService.Update(ctx, userData.ID, payload, userData.TenantID)
	if err != nil {
		return
	}

	// Delete reset token
	payload = &user.UserUpdate{
		ResetToken:     func() *string { s := ""; return &s }(),
		ResetExpiresAt: nil,
	}
	_, err = s.userService.Update(ctx, userData.ID, payload, userData.TenantID)
	if err != nil {
		return
	}

	// TODO: Send email

	return s.Login(ctx, updateduser.Email, newPassword)
}

func (s *serviceImpl) RequestPasswordReset(ctx context.Context, email string) error {
	// Find user by email
	u, err := s.userService.FindByEmail(ctx, email)
	if err != nil {
		return err
	}

	// Generate reset token
	token := utils.GenerateSecureToken()

	expiration := time.Now().Add(time.Minute * 5)
	payload := &user.UserUpdate{
		ResetToken:     &token,
		ResetExpiresAt: &expiration,
	}

	fmt.Println("reset token:", token)

	// Call user service to update credentials
	if _, err := s.userService.Update(ctx, u.ID, payload, u.TenantID); err != nil {
		return err
	}

	// TODO: Send email

	return nil
}

func (s *serviceImpl) generateJWT(u *user.User) (tokenStr string, err error) {
	builder := jwt.NewBuilder().
		Issuer(s.cfg.JWT.Issuer).
		Subject(u.ID.String()).
		IssuedAt(time.Now()).
		Expiration(time.Now().Add(time.Second * time.Duration(s.cfg.JWT.ExpirySeconds)))

	if u.TenantID != nil {
		builder.Claim("tenant_id", u.TenantID.String())
	}

	token, err := builder.Build()
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
	builder := jwt.NewBuilder().
		Issuer(s.cfg.JWT.Issuer).
		Subject(u.ID.String()).
		IssuedAt(time.Now()).
		Expiration(time.Now().Add(time.Second * time.Duration(s.cfg.JWT.ExpirySeconds)))

	if u.TenantID != nil {
		builder.Claim("tenant_id", u.TenantID.String())
	}

	refreshToken, err := builder.Build()
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

func (s *serviceImpl) validateResetToken(ctx context.Context, token string) (*user.User, error) {
	// Find user by reset token
	u, err := s.userService.FindByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	// Check expiration
	if u.ResetExpiresAt == nil || time.Now().After(*u.ResetExpiresAt) {
		return nil, fmt.Errorf("token expired")
	}

	// Token is valid
	return u, nil
}
