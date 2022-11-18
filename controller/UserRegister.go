package controller

import (
	"Wave/database"
	"Wave/model"
	"Wave/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserRegisterInfo struct {
	UserName string `json:"userName" form:"userName" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Email    string `json:"email" form:"email" binding:"email,required"`
}

func UserRegister(c *gin.Context) {
	var u UserRegisterInfo
	if err := c.ShouldBind(&u); err != nil {
		fmt.Println("参数校验不通过,", err)
		c.JSON(http.StatusBadRequest, &util.Restful{
			Code:    0,
			Message: "参数校验不通过," + err.Error(),
			Data:    nil,
		})
		return
	}
	passwordDigest, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 12)

	newUser := &model.User{
		UserName: u.UserName,
		Password: string(passwordDigest),
		Email:    u.Email,
	}
	res := database.DB.Create(&newUser)
	if res.Error != nil {
		println("创建失败,", res.Error.Error())
		c.JSON(http.StatusOK, &util.Restful{
			Code:    0,
			Message: "创建失败," + res.Error.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, &util.Restful{
		Code:    0,
		Message: "创建成功",
		Data:    nil,
	})
}
