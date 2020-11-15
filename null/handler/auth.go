package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/dao"
)

//账号注册
func Signup(c *gin.Context) {
	var user = &dao.AuthUser{}
	err := c.BindJSON(user)
	if err != nil {
		log.Print("func signup: binding json err:", err)
	}
	err = user.StoreUser()
	if err != nil {
		log.Print(err)
	}
}

//账号登陆
func Signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg": "账号登陆路由",
	})
}
