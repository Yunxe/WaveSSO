package sso

import (
	"Wave/util"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) (err error, data any) {
	var (
		h   util.Header
		arr []string
	)
	if err := c.ShouldBindHeader(&h); err != nil {
		c.Abort()
		return util.AUTH_REQUIRE, nil
	}

	if arr = strings.Fields(h.Authorization); strings.ToLower(arr[0]) != "bearer" {
		c.Abort()
		return util.AUTH_TYPE_ERR, nil
	}

	token, err := util.ParseToken(arr[1])
	if err != nil {
		c.Abort()
		switch {
		case errors.Is(err, util.AUTH_TOKEN_EXPIRED):
			return util.AUTH_TOKEN_EXPIRED, nil
		case errors.Is(err, util.AUTH_TOKEN_INVALID_ISSUER):
			return util.AUTH_TOKEN_INVALID_ISSUER, nil
		default:
			return util.AUTH_PARSE_TOKEN_ERR, nil
		}
	}
	if claims, ok := token.Claims.(*util.UserClaims); ok && token.Valid {
		c.Set("uid", claims.Uid)
		c.Next()
	}
	u, _ := c.Get("user")
	return nil, &util.StatusWithData{
		Code:    0,
		Message: "成功",
		Data:    u,
	}

}
