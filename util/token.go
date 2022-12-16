package util

import (
	"Wave/config"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenInfo struct {
	Token     string
	TokenType string
	ExpiresIn time.Duration
}

type UserClaims struct {
	Uid    uint `json:"uid"`
	Role   int8 `json:"role"`
	Status int8 `json:"status"`
	jwt.RegisteredClaims
}

type Header struct {
	Authorization string `header:"Authorization" binding:"required"`
}

func (c *UserClaims) CreateToken() (string, error) {
	SignKey := os.Getenv("SIGNING-KEY")

	c.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "WaveSSO",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.ExpireTime)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedToken, err := token.SignedString([]byte(SignKey))
	return signedToken, err
}

func ParseToken(unauthToken string) (*jwt.Token, error) {
	SignKey := os.Getenv("SIGNING-KEY")

	return jwt.ParseWithClaims(unauthToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
}
