package main

import (
	"bobo/config"
	"bobo/middlewares"
	"bobo/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	config := config.Get()
	gin.SetMode(config.RunMode)
	r := gin.Default()

	/// 中间件
	r.Use(middlewares.Cors)
	r.Use(middlewares.Recover)
	/// routers
	routers.InitRouter(r)
	// 开启服务
	r.Run(":" + config.Server.Port)
}
