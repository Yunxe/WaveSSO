package sso

import (
	"Wave/database"
	"Wave/model"
	"Wave/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserRegisterInfo struct {
	UserName string `json:"userName" form:"userName" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Email    string `json:"email" form:"email" binding:"email,required"`
}

func UserRegister(c *gin.Context) (err error, data any) {
	var u UserRegisterInfo
	if err := c.ShouldBind(&u); err != nil {
		return util.REQ_PARAM_INVALID_ERR, nil
	}

	passwordDigest, err := bcrypt.GenerateFromPassword([]byte(u.Password), 16)
	if err != nil {
		return err, nil
	}

	newUser := &model.User{
		UserName: u.UserName,
		Password: string(passwordDigest),
		Email:    u.Email,
	}
	res := database.DB.Create(&newUser)
	if res.Error != nil {
		return res.Error, nil
	}

	//c.Redirect(http.StatusNotModified,"/api/login")
	return nil, &util.SUCCESS.Status
}
