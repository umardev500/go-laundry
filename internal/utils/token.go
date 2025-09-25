package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateSecureToken return a URL-safe, base64 encoded random string
func GenerateSecureToken() string {
	b := make([]byte, 32) // 32 bytes = 256 bit
	if _, err := rand.Read(b); err != nil {
		return ""
	}

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)
}
