package sso

import (
	"Wave/database"
	"Wave/util"
	"errors"
	"github.com/go-redis/redis/v9"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) {
	var (
		h   util.Header
		arr []string
	)
	if err := c.ShouldBindHeader(&h); err != nil {
		c.AbortWithStatusJSON(
			util.ErrWrapper(util.AUTH_REQUIRE),
		)
	}

	if arr = strings.Fields(h.Authorization); strings.ToLower(arr[0]) != "bearer" {
		c.AbortWithStatusJSON(
			util.ErrWrapper(util.AUTH_TYPE_ERR),
		)
	}

	token, err := util.ParseToken(arr[1])
	if err != nil {
		switch {
		case errors.Is(err, util.AUTH_TOKEN_EXPIRED):
			c.AbortWithStatusJSON(
				util.ErrWrapper(util.AUTH_TOKEN_EXPIRED),
			)
		case errors.Is(err, util.AUTH_TOKEN_INVALID_ISSUER):
			c.AbortWithStatusJSON(
				util.ErrWrapper(util.AUTH_TOKEN_INVALID_ISSUER),
			)
		default:
			c.AbortWithStatusJSON(
				util.ErrWrapper(util.AUTH_PARSE_TOKEN_ERR),
			)
		}
	}
	if claims, ok := token.Claims.(*util.UserClaims); ok && token.Valid {
		err := database.RDB.Get(c, token.Raw).Err()
		if err != redis.Nil {
			c.AbortWithStatusJSON(
				util.ErrWrapper(util.AUTH_TOKEN_INVALID_IN_BLACK_LIST),
			)
		}

		c.Set("uid", claims.Uid)
		c.Set("token", token.Raw)

		c.Next()
	}

}
