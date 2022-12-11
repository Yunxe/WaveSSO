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
		v1.POST("/register", util.HandlerWarpper(sso.UserRegister))
		v1.POST("/login", util.HandlerWarpper(sso.UserLogin))
		v1.POST("/info", util.HandlerWarpper(sso.UserAuth), sso.UserInfo)
	}
	if err := r.Run(":8079"); err != nil {
		fmt.Println(err)
		return
	}
}
