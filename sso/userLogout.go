package sso

import (
	"Wave/config"
	"Wave/database"
	"Wave/model"
	"Wave/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserLogout(c *gin.Context) (err error, data any) {
	user := model.NewUser()

	uid, _ := c.Get("uid")
	token, _ := c.Get("token")

	database.DB.Where("uid = ?", uid).First(&user)
	if *user == *model.NewUser() {
		return util.USER_NOT_FOUND, nil
	}

	str := fmt.Sprintf("%v", token)

	err = database.RDB.Set(c, str, uid, config.ExpireTime).Err()
	if err != nil {
		return err, nil
	}

	return nil, &util.SUCCESS.Status
}
