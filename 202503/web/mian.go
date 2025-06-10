package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的 Gin 路由引擎
	r := gin.Default()

	// 定义一个 GET 路由，路径为 /hello
	r.GET("/hello", func(c *gin.Context) {
		// 返回一个 JSON 响应
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// 启动服务，默认监听 0.0.0.0:8080
	r.Run("0.0.0.0:9090")
}
