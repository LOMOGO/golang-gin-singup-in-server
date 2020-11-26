//用于用户登陆注册的handler函数
package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"server/dao"
	"server/extend/standardCode"
	"server/extend/utils"
	"server/extend/validata"
)

//账号注册
func Signup(c *gin.Context) {
	var user = dao.SignupUser{}
	err := c.BindJSON(&user)
	//因为gin中集成了validata，因此当接口中的参数不符合结构体中的banding条件时就会报错
	if err != nil {
		//利用validata库进行接口值的校验
		errs := err.(validator.ValidationErrors)
		standardCode.ValidataError.Message = errs.Translate(validata.Trans)
		standardCode.CodeFormatter(c, standardCode.ValidataError, nil)
		return
	}

	//在数据库中搜索用户名是否存在
	_, err = dao.FindUserByName(user.Name)
	if err == nil {
		//标准化状态码
		standardCode.CodeFormatter(c, standardCode.RepeatedUserError, nil)
		return
	}
	//将密码进行哈希加密
	pwd := utils.MakeSha1(user.Password)
	user.Password = pwd
	//如果用户不存在，那么就存入数据库
	err = user.StoreUser()
	if err != nil {
		standardCode.CodeFormatter(c, standardCode.InternalServerError, nil)
		return
	}
	standardCode.CodeFormatter(c, standardCode.Success, nil)
}

//账号登陆
func Signin(c *gin.Context) {
	var user = &dao.SigninUser{}
	err := c.BindJSON(user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		standardCode.ValidataError.Message = errs.Translate(validata.Trans)
		standardCode.CodeFormatter(c, standardCode.ValidataError, nil)
		return
	}
	//在数据库中搜寻用户名是否存在
	hashPwd, err := dao.FindUserByName(user.Name)
	if err != nil {
		//用户名或密码错误
		standardCode.CodeFormatter(c, standardCode.IllegalNameOrPwdError, nil)
		return
	}
	//比较哈希密码与明文密码是否一致
	if hashPwd == utils.MakeSha1(user.Password) {
		standardCode.CodeFormatter(c, standardCode.Success, nil)
	} else {
		//用户名或密码错误
		standardCode.CodeFormatter(c, standardCode.IllegalNameOrPwdError, nil)
	}

}
