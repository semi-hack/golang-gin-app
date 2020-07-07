package services

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)


var key = []byte(os.Getenv("SECRET_KEY"))

// Claims ...
type Claims struct {
	Username string `json:"name"`
	jwt.StandardClaims
}

// GenerateToken ...
func GenerateToken(Username string) (string, error) {
	expirationTime := time.Now().Add(1440 *time.Minute)

	// set claims i.e (payload)
	claims := &Claims{
		Username,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//sign token with secret key
	ts, err := token.SignedString(key)

	return ts, err
}

// func verifyToken(tokenString string) (*jwt.Token, string) {
// 	claims := &Claims{}

// 	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

// 	})
// }