package sso

import (
	"Wave/database"
	"Wave/model"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	var (
		user model.User
	)

	uid, _ := c.Get("uid")

	database.DB.Where("uid = ?", uid).First(&user)
	c.Set("user", user)
}
