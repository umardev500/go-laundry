package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/redis/go-redis/v9"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/domain/auth"
	platformuser "github.com/umardev500/go-laundry/internal/domain/platform_user"
	tenantuser "github.com/umardev500/go-laundry/internal/domain/tenant_user"
	"github.com/umardev500/go-laundry/internal/domain/user"
	"github.com/umardev500/go-laundry/internal/types"
	"github.com/umardev500/go-laundry/internal/utils"
	"github.com/umardev500/go-laundry/pkg/email"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(ctx context.Context, email, password string) (user *user.User, token, refreshToken string, err error)
	ResetPassword(ctx context.Context, token, newPassword string, scope *types.Scoped) (user *user.User, accessToken, refreshToken string, err error)
	RequestPasswordReset(ctx context.Context, email string, scope *types.Scoped) error
}

type serviceImpl struct {
	userService     user.Service
	cfg             *config.Config
	emailClient     *email.EmailClient
	repo            auth.Repository
	platformUserSrv platformuser.Service
	tenantUserSrv   tenantuser.Service
}

func NewServiceImpl(
	userService user.Service,
	cfg *config.Config,
	emailClient *email.EmailClient,
	repo auth.Repository,
	platformUserSrv platformuser.Service,
	tenantUserSrv tenantuser.Service,
) *serviceImpl {
	return &serviceImpl{
		cfg:             cfg,
		repo:            repo,
		userService:     userService,
		emailClient:     emailClient,
		platformUserSrv: platformUserSrv,
		tenantUserSrv:   tenantUserSrv,
	}
}

func (s *serviceImpl) Login(ctx context.Context, email, password string) (user *user.User, token, refreshToken string, err error) {
	user, err = s.userService.FindByEmail(ctx, email)
	if err != nil {
		return
	}

	platformUser, err := s.platformUserSrv.GetByUserID(ctx, user.ID)
	if err != nil {
		return
	}

	fmt.Println(platformUser)

	tenantUser, err := s.tenantUserSrv.GetByUserID(ctx, user.ID)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Tenant user", tenantUser)

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return
	}

	var tenantID = uuid.Nil
	if user.TenantID != nil {
		tenantID = *user.TenantID
	}

	planID, err := s.repo.GetActivePlanID(ctx, tenantID)
	if err != nil && err != redis.Nil {
		return
	}

	token, err = s.generateJWT(user, planID)
	if err != nil {
		return
	}

	refreshToken, err = s.generateRefreshToken(user, planID)
	if err != nil {
		return
	}

	return
}

func (s *serviceImpl) ResetPassword(ctx context.Context, token, newPassword string, scope *types.Scoped) (u *user.User, accessToken, refreshToken string, err error) {
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
	updateduser, err := s.userService.Update(ctx, userData.ID, payload, scope)
	if err != nil {
		return
	}

	// Delete reset token
	payload = &user.UserUpdate{
		ResetToken:     func() *string { s := ""; return &s }(),
		ResetExpiresAt: nil,
	}
	_, err = s.userService.Update(ctx, userData.ID, payload, scope)
	if err != nil {
		return
	}

	// Send message to the user email
	go func() {
		s.emailClient.Send([]string{updateduser.Email}, "Password Reset", "Your password has been reset")
	}()

	return s.Login(ctx, updateduser.Email, newPassword)
}

func (s *serviceImpl) RequestPasswordReset(ctx context.Context, email string, scope *types.Scoped) error {
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

	// Call user service to update credentials
	if _, err := s.userService.Update(ctx, u.ID, payload, scope); err != nil {
		return err
	}

	// TODO: Send email

	// Send message to the user email
	go func() {
		s.emailClient.Send(
			[]string{u.Email},
			"Password Reset",
			"Please reset your password using the following link: http://localhost:8080/reset-password?token="+token,
		)
	}()

	return nil
}

func (s *serviceImpl) generateJWT(u *user.User, planID uuid.UUID) (tokenStr string, err error) {
	builder := jwt.NewBuilder().
		Issuer(s.cfg.JWT.Issuer).
		Subject(u.ID.String()).
		IssuedAt(time.Now()).
		Expiration(time.Now().Add(time.Second * time.Duration(s.cfg.JWT.ExpirySeconds)))

	if u.TenantID != nil {
		builder.Claim("tenant_id", u.TenantID.String())
	}

	if planID != uuid.Nil {
		builder.Claim("plan_id", planID.String())
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

func (s *serviceImpl) generateRefreshToken(u *user.User, planID uuid.UUID) (tokenStr string, err error) {
	builder := jwt.NewBuilder().
		Issuer(s.cfg.JWT.Issuer).
		Subject(u.ID.String()).
		IssuedAt(time.Now()).
		Expiration(time.Now().Add(time.Second * time.Duration(s.cfg.JWT.ExpirySeconds)))

	if u.TenantID != nil {
		builder.Claim("tenant_id", u.TenantID.String())
	}

	if planID != uuid.Nil {
		builder.Claim("plan_id", planID.String())
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
