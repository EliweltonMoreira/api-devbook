package security

import "golang.org/x/crypto/bcrypt"

// Hash receive a string and put a hash on her
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword compares a password and a hash and returns if they match
func VerifyPassword(passwordWithHash, stringPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(stringPassword))
}
