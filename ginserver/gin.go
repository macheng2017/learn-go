package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const keyRequestId = "requestId"

func main() {
	r := gin.Default()
	// 添加一个中间件
	logger, e := zap.NewProduction()
	if e != nil {
		panic(e)
	}

	r.Use(func(c *gin.Context) {
		s := time.Now()
		// 洋葱模型?
		c.Next()
		// 结构化log信息
		// path , log latency, response code
		logger.Info("incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()), zap.Duration("elapsed", time.Now().Sub(s)))

	},
		// 再添加一个中间件,在context中添加一个值
		func(c *gin.Context) {

			c.Set(keyRequestId, rand.Int())
		})

	r.GET("/ping", func(c *gin.Context) {
		rid, exists := c.Get(keyRequestId)

		hs := gin.H{
			"message": "pong",
		}
		if exists {
			hs[keyRequestId] = rid
		}
		c.JSON(200, hs)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "world",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
