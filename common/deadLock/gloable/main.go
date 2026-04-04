package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // 引入 pprof 工具
	"sync"
	"time"
)

func main() {
	// 1. 开启 pprof 监控服务（必须在独立 Goroutine 中）
	go func() {
		fmt.Println("pprof 服务器已启动：http://localhost:6060/debug/pprof/goroutine?debug=2")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			fmt.Println("pprof 启动失败:", err)
		}
	}()

	var muA, muB sync.Mutex

	// 2. 制造两个协程的局部死锁（相互等待）
	go func() { // Goroutine 1
		muA.Lock()
		time.Sleep(100 * time.Millisecond) // 模拟业务处理
		muB.Lock()                         // 等待 muB 释放
		muB.Unlock()
		muA.Unlock()
	}()

	go func() { // Goroutine 2
		muB.Lock()
		time.Sleep(100 * time.Millisecond) // 模拟业务处理
		muA.Lock()                         // 等待 muA 释放
		muA.Unlock()
		muB.Unlock()
	}()

	// 3. 主协程持续运行（避免全局死锁被自动检测到）
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("主协程还在运行，但那两个 Goroutine 已经死锁了...")
	}
}
