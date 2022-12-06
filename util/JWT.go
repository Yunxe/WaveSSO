package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type UserClaims struct {
	Uid    uint `json:"uid"`
	Role   int8 `json:"role"`
	Status int8 `json:"status"`
	jwt.RegisteredClaims
}

func (c *UserClaims) CreateToken() (string, error) {
	signKey := os.Getenv("SIGNING-KEY")

	c.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "Wave",
		Subject:   "SSO",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}
	fmt.Println(signKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	ss, err := token.SignedString([]byte(signKey))
	return ss, err
}
