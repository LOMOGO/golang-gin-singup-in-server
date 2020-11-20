//路由
package router

import (
	"github.com/gin-gonic/gin"
	v12 "server/handler/v1"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	//登陆注册路由分组
	Authenticator := v1.Group("/auth")
	{
		//账号注册路由
		Authenticator.POST("/signup", v12.Signup)
		//账号登陆路由
		Authenticator.POST("/signin", v12.Signin)
	}

	return r
}
