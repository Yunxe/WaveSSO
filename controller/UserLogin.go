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

type UserLoginInfo struct {
	Email    string `json:"email" form:"email" binding:"email,required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func UserLogin(c *gin.Context) {
	var (
		u    UserLoginInfo
		user model.User
	)
	if err := c.ShouldBind(&u); err != nil {
		fmt.Println("参数校验不通过,", err)
		c.JSON(http.StatusBadRequest, &util.Restful{
			Code:    0,
			Message: "参数校验不通过," + err.Error(),
			Data:    nil,
		})
		return
	}

	database.DB.Where("email = ?", u.Email).First(&user)

	if user == *model.NilUser {
		c.JSON(http.StatusBadRequest, &util.Restful{
			Code:    0,
			Message: "用户不存在",
			Data:    nil,
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		c.JSON(http.StatusBadRequest, &util.Restful{
			Code:    0,
			Message: "密码错误",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, &util.Restful{
		Code:    1,
		Message: "登录成功",
		Data:    u,
	})
}
