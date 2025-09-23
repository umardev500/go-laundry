package httputil

import (
	"errors"
	"strings"
)

// ExtractBearerToken parses "Bearer <token>" from Authorization header
func ExtractBearerToken(header string) (string, error) {
	parts := strings.Split(header, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid authorization header format")
	}
	return parts[1], nil
}
