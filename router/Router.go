package router

import (
	"Wave/sso"
	"Wave/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewRouter() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api")
	{
		v1.POST("/register", util.HandlerWarpper(sso.UserRegister))
		v1.POST("/login", util.HandlerWarpper(sso.UserLogin))
	}
	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
		return
	}
}
