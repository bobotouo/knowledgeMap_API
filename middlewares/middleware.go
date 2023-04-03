package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// / 错误处理
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  errorToString(r),
				"data": nil,
			})
			c.Abort()
		}
	}()
	c.Next()
}

func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

// // 跨域
func Cors(context *gin.Context) {
	method := context.Request.Method
	context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	context.Header("Access-Control-Allow-Credentials", "true")
	if method == "OPTIONS" {
		context.AbortWithStatus(http.StatusNoContent)
	}
	context.Next()
}
