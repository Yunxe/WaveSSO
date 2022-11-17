package router

import (
	"Wave/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(){
	r:=gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	
	v1:=r.Group("/api")
	{
		v1.POST("/register",controller.UserRegister)
	}
	r.Run(":8080")
}