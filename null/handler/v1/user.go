//用于操作用户信息的handler函数
package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"server/dao"
	"server/extend/standardCode"
	"server/extend/token"
	"server/extend/utils"
	"server/extend/validata"
)

//获取用户信息
func GetUserInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*token.CustomClaims)
	id := claims.UserID
	name, gender, createTime, err := dao.UseIDGetInfo(id)
	if err != nil {
		standardCode.CodeFormatter(c, standardCode.FailedGetUserInfo, nil)
		return
	}
	standardCode.CodeFormatter(c, standardCode.Success, gin.H{
		"username": name,
		"gender": gender,
		"createTime": createTime,
	})
}

type AlterPwd struct {
	OldPassword string `json:"old_password" binding:"required,max=16,min=8" label:"旧密码"`
	NewPassword string `json:"new_password" binding:"required,max=16,min=8,nefield=OldPassword" label:"新密码"`
}

//修改用户密码
func AlterUserPwd(c *gin.Context) {
	alterPwd := AlterPwd{}
	err := c.BindJSON(&alterPwd)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		standardCode.ValidataError.Message = errs.Translate(validata.Trans)
		standardCode.CodeFormatter(c, standardCode.ValidataError, nil)
		return
	}
	claims := c.MustGet("claims").(*token.CustomClaims)
	username := claims.Username
	pwd, id, _ := dao.FindUserByName(username)
	userid := id.(uint)
	if userid == claims.UserID && utils.MakeSha1(alterPwd.OldPassword) == pwd {
		err := dao.UpdateUserPwd(userid, alterPwd.NewPassword)
		if err != nil {
			log.Println("UpdatePwd:未能更新密码", err)
			standardCode.CodeFormatter(c, standardCode.FailedUpdatePwd, nil)
			return
		}
		standardCode.CodeFormatter(c, standardCode.Success, gin.H{
			"newPwd": alterPwd.NewPassword,
		})
	}
}
