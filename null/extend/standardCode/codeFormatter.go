package standardCode

import "github.com/gin-gonic/gin"

//将业务码标准化
func CodeFormatter(c *gin.Context, code Code, data interface{}) {
	c.JSON(code.Status, gin.H{
		"code": code.Code,
		"message": code.Message,
		"data": data,
	})
}
