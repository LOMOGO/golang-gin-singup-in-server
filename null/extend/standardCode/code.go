//用于定义标准码
package standardCode

import "net/http"

//标准码
type Code struct {
	Status int `json:"statu"`
	Code int `json:"code"`
	Message interface{} `json:"message"`
}

var (
	//请求成功
	Success = Code{Status: http.StatusOK, Code: 2000001, Message: "请求成功🤗"}
	//服务器内部错误
	InternalServerError = Code{Status: http.StatusInternalServerError, Code: 5000000, Message: "服务器内部错误🙁"}
	//用户名或密码错误
	IllegalNameOrPwdError = Code{Status: http.StatusBadRequest, Code: 4000001, Message: "用户名或密码错误🙁"}
	//用户名已存在
	RepeatedUserError = Code{Status: http.StatusBadRequest, Code: 4000002, Message: "用户名已存在，换个试试吧🙁"}
	//格式验证错误
	ValidataError = Code{Status: http.StatusBadRequest, Code: 4000003}
	//Token无效
	TokenInvalid = Code{Status: http.StatusBadRequest, Code: 4000004, Message: "无效的token，拒绝访问🙁"}
	//请求未携带token
	TokenNotFound = Code{Status: http.StatusBadRequest, Code: 4000005, Message: "请求头中未发现token信息，拒绝访问🙁"}
	//token未生成
	TokenNotGenerate = Code{Status: http.StatusInternalServerError, Code: 5000001, Message: "token生成失败🙁"}
	//获取用户信息失败
	FailedGetUserInfo = Code{Status: http.StatusBadRequest, Code: 4000006, Message: "获取用户信息失败"}
	//密码更新失败
	FailedUpdatePwd = Code{Status: http.StatusBadRequest, Code: 4000007, Message: "密码更新失败"}
)
