package middleware

import (
	"github.com/gin-gonic/gin"
	"server/extend/standardCode"
	"server/extend/token"
)

//验证token是否有效的中间件
func TokenChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		//c.Abort方法调用后仍然需要使用return函数，因为abort虽然可以终止后续函数的调用，但是不会终止当前处理函数中的其他语句
		if tokenString == "" {
			standardCode.CodeFormatter(c, standardCode.TokenNotFound, nil)
			c.Abort()
			return
		}

		claims, ok := token.ParseToken(tokenString)
		if !ok {
			standardCode.CodeFormatter(c, standardCode.TokenInvalid, nil)
			c.Abort()
			return
		}

		//将claim传递给下一个handler函数
		c.Set("claims", claims)
		c.Next()
	}
}
