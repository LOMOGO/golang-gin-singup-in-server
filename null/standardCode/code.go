package standardCode

import "net/http"

//标准码
type Code struct {
	Status int `json:"statu"`
	Code int `json:"code"`
	Message string `json:"message"`
}

var (
	//请求成功
	Success = Code{Status: http.StatusOK, Code: 2000001, Message: "请求成功😙"}
	//服务器内部错误
	InternalServerError = Code{Status: http.StatusInternalServerError, Code: 5000000, Message: "服务器内部错误🙁"}
)

