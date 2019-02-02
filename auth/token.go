package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
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
	tokenString, err := tokenJwt.SigningString()
	if err != nil {
		logrus.Error("can't create auth token: ", err)
		return nil
	}
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

// IsValid .. Check if the token is valid
func (t *Token) IsValid() {

}