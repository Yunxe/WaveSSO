package sso

import (
	"Wave/database"
	"Wave/model"
	"Wave/util"
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) (err error, data any) {
	user := model.NewUser()

	uid, _ := c.Get("uid")

	database.DB.Where("uid = ?", uid).First(&user)
	if *user == *model.NewUser() {
		return util.USER_NOT_FOUND, nil
	}

	return nil, util.StatusWithData{
		Code:    0,
		Message: "成功",
		Data:    user,
	}
}
