package security

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

// Checks if the password given matches the hashed password
//
// Arguments:
//   - password	The password to verify
//   - hash: 		The hash of the password saved in DB
func CheckPasswordHash(password string, hash string) bool {
	matched := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return matched == nil
}
