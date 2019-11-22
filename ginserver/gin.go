package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	r := gin.Default()
	// 添加一个中间件
	r.Use(func(c *gin.Context) {
		// 结构化log信息
		// path , log latency, response code
		logger, e := zap.NewProduction()
		if e != nil {
			panic(e)
		}
		logger.Info("incoming request", zap.String("path", c.Request.URL.Path))
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "world",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
