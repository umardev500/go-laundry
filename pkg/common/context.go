package common

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const ClaimsKey ContextKey = "claims"

func GetClaims[T jwt.Claims](ctx context.Context) (T, bool) {
	var claims T

	if claims, ok := ctx.Value(ClaimsKey).(T); ok {
		return claims, true
	}

	return claims, false
}
