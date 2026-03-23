package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"outback/algorithm/common/RateLimiter/single/leakybucket"
)

// leakyBucketMiddleware 返回一个基于漏桶管理器的 Gin 中间件。
func leakyBucketMiddleware(mgr *leakybucket.Manager, keyFunc func(c *gin.Context) string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := keyFunc(c)
		if key == "" {
			key = "anonymous"
		}
		if !mgr.Allow(key) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code":    "TOO_MANY_REQUESTS",
				"message": "rate limit exceeded",
				"key":     key,
			})
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	if err := r.SetTrustedProxies(nil); err != nil {
		log.Fatalf("failed to set trusted proxies: %v", err)
	}

	// 每个 key 独立限流：每秒漏出 5 个请求，桶容量 10。
	limiterManager := leakybucket.NewManager(5, 10)

	r.Use(leakyBucketMiddleware(limiterManager, func(c *gin.Context) string {
		// 这里使用客户端 IP 作为限流 key，实战中也可以换成 userID、apiKey 等。
		return c.ClientIP()
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"ts":      time.Now().UnixMilli(),
		})
	})

	r.GET("/work", func(c *gin.Context) {
		time.Sleep(120 * time.Millisecond)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	addr := ":8080"
	log.Printf("server starting at %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server start failed: %v", err)
	}
}
