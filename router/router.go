package router

import (
	"Wave/sso"
	"Wave/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewRouter() {
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api")
	{
		v1.POST("/register", util.HandlerWrapper(sso.UserRegister))
		v1.POST("/login", util.HandlerWrapper(sso.UserLogin))
		v1.POST("/info", sso.UserAuth, util.HandlerWrapper(sso.UserInfo))
		v1.POST("/logout", sso.UserAuth, util.HandlerWrapper(sso.UserLogout))
	}
	if err := r.Run(":8079"); err != nil {
		fmt.Println(err)
		return
	}
}
