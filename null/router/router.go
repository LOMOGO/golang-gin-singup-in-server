//路由
package router

import (
	"github.com/gin-gonic/gin"
	v12 "server/handler/v1"
	"server/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	//登陆注册路由分组
	authenticator := v1.Group("/auth")
	{
		//账号注册路由
		authenticator.POST("/signup", v12.Signup)
		//账号登陆路由
		authenticator.POST("/signin", v12.Signin)
	}
	//用户信息操作路由信息分组
	user := v1.Group("/user")
	user.Use(middleware.TokenChecker())
	{
		//获取用户信息的路由
		user.POST("/getInfo", v12.GetUserInfo)
		//更改用户密码的路由
		user.PATCH("/alterPwd", v12.AlterUserPwd)
	}

	return r
}
