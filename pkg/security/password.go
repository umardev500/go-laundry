package security

import "golang.org/x/crypto/bcrypt"

var DefaultCost = bcrypt.DefaultCost

// HashPassword hashes the given plain password using bcrypt.
// Returns the hashed password as string.
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	return string(hashed), err
}

// CheckPasswordHash checks if the given password matches the hash.
// Returns true if the password matches the hash, false otherwise.
func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
