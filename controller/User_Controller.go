package controller

import (
	"Wave/database"
	"Wave/model"
	"Wave/service"
	"Wave/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var u service.UserRegisterInfo
	if err :=c.ShouldBind(&u); err != nil {
		fmt.Println("参数校验不通过", err)
		c.JSON(http.StatusBadRequest, &util.Restful{
			Code:    0,
			Message: "参数校验不通过"+err.Error(),
			Data:    nil,
		})
		return
	}

	newUser := &model.User{
		UserName: u.UserName,
		Password: u.Password,
		Email: u.Email,
	}
	res := database.DB.Create(&newUser)
	if res.Error != nil {
		println("创建失败", res.Error.Error())
		c.JSON(http.StatusOK,&util.Restful{
			Code:    0,
			Message: "创建失败"+res.Error.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK,&util.Restful{
		Code:    0,
		Message: "创建成功",
		Data:    nil,
	})
}
