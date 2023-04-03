package routers

import (
	"bobo/services"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	r.GET("/all", services.AllHandlerFunc())
	r.GET("/query", services.QueryHandlerFunc())
	r.GET("/search", services.SearchHandleFunc())
}
