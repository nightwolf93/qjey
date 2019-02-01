package auth

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type Token struct {
	TokenString string
	CreatedAt time.Time
	ExpireAt time.Time
}

// NewToken .. create a new token
func NewToken() *Token {
	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().UTC().Unix(),
		"exp": time.Now().UTC().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, _ := tokenJwt.SignedString(os.Getenv("JWT_SIGNING_KEY"))
	token := &Token{
		TokenString: tokenString,
		CreatedAt: time.Now().UTC(),
		ExpireAt: time.Now().UTC().Add(time.Hour * 24 * 30), // This token is valid for 30 days
	}
	return token
}

// ParseToken .. parse a row token string into the token object
func ParseToken(tokenString string) {

}

// Check if the token is valid
func (t *Token) IsValid() {

}