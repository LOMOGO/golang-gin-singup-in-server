package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/dao"
	"server/standardCode"
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
		standardCode.CodeFormatter(c, standardCode.InternalServerError, nil)
		return
	}
	standardCode.CodeFormatter(c, standardCode.Success, nil)
}

//账号登陆
func Signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg": "账号登陆路由",
	})
}
