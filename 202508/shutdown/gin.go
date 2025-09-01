package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 创建 Gin 路由
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		// 模拟一个耗时5秒的处理
		log.Println("Request to /test started...")
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Request finished after 5 seconds")
		log.Println("Request to /test finished.")
	})

	// 2. 创建 http.Server
	// 我们不直接使用 router.Run(":8080")
	// 而是创建一个 http.Server，这样我们能获得对服务器生命周期的完全控制
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 3. 在一个新的 goroutine 中启动服务器
	// ListenAndServe 是一个阻塞操作，所以需要放在 goroutine 中
	// 这样主 goroutine 就不会被阻塞，可以继续执行后面的信号监听逻辑
	go func() {
		log.Println("Server is starting and listening on port 8080...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err)
		}
	}()

	// 4. 设置信号监听和优雅退出逻辑
	// 创建一个 channel 用于接收系统信号
	quit := make(chan os.Signal, 1)
	// signal.Notify 会将指定的信号转发到 quit channel
	// 我们监听 SIGINT (Ctrl+C) 和 SIGTERM (kill 命令)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞主 goroutine，直到接收到信号
	<-quit
	log.Println("Shutdown signal received, shutting down server...")

	// 创建一个有5秒超时的 context
	// 这给了正在处理的请求5秒钟的时间来完成
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 调用 srv.Shutdown() 来优雅地关闭服务器
	// 这个方法会阻塞，直到所有打开的连接都已关闭或超时
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting gracefully.")
}
