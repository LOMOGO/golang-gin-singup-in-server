package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/handler"
)

func Setup() {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	//登陆注册路由分组
	Authenticator := v1.Group("/auth")
	{
		//账号注册路由
		Authenticator.POST("/signup", handler.Signup)
		//账号登陆路由
		Authenticator.POST("/signin", handler.Signin)
	}

	err := r.Run(":8080")
	if err != nil {
		log.Print(err)
	}
}
