package security

import (
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GetJWTSecret() string {
	godotenv.Load()
	return os.Getenv("JWT_SECRET")
}

func VerifyJWT(tokenToVerify string) (*jwt.Token, error) {
	jwtSecret := GetJWTSecret()

	return jwt.Parse(tokenToVerify, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(jwtSecret), nil
	})

}
