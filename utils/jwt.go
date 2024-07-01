package utils

import (
	// "crypto/rand"
	// "encoding/base64"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECERET")))
	if err != nil {
		panic(err)
	}
	return tokenString, err
}
