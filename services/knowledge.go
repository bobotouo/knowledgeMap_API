package services

import (
	"bobo/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllHandlerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		eid, b := ctx.GetQuery("eid")
		if b {
			p := repositories.GetAllData(eid)
			ctx.JSON(200, gin.H{
				"data": p,
				"code": 200,
			})
		} else {
			ctx.JSON(400, gin.H{
				"data": nil,
				"code": 400,
				"msg":  "缺少参数",
			})
		}
	}
}

func QueryHandlerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		eid := ctx.DefaultQuery("eid", "0")
		p := repositories.QueryNode(eid)
		ctx.JSON(200, gin.H{
			"data": p,
			"code": 200,
		})
	}
}

func SearchHandleFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		content := ctx.DefaultQuery("text", "")
		lists := repositories.SearchContent(content)
		ctx.JSON(http.StatusOK, gin.H{
			"data": lists,
			"code": 200,
		})
	}
}
