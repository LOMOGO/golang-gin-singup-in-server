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
	IllegalNameOrPwdError = Code{Status: http.StatusForbidden, Code: 4000003, Message: "用户名或密码错误🙁"}
	//用户名已存在
	RepeatedUserError = Code{Status: http.StatusForbidden, Code: 4000003, Message: "用户名已存在，换个试试吧🙁"}
	//格式验证错误
	ValidataError = Code{Status: http.StatusForbidden, Code: 4000003}
)

//数据验证错误信息
/*func ValidataErrorMsg() string {

}*/

