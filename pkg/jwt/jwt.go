package jwt

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	sharedctx "github.com/umardev500/go-laundry/pkg/context"
)

func Sign(claims jwt.Claims, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func Parse(tokenString string, claims jwt.Claims, secret []byte) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return secret, nil
	})
}

func Claims[T jwt.Claims](ctx context.Context) (T, error) {
	var claims T

	claimsCtx := ctx.Value(sharedctx.ClaimsContextKey)
	if claimsCtx == nil {
		return claims, sharedctx.ErrNotFound
	}

	claims, ok := claimsCtx.(T)
	if !ok {
		return claims, sharedctx.ErrNotFound
	}

	return claims, nil
}
