// 实验前需要先安装： go get github.com/sasha-s/go-deadlock
package main

import (
	"fmt"
	"time"

	"github.com/sasha-s/go-deadlock"
)

func main() {
	// 1. 设置死锁超时时间（如果在指定时间内获取不到锁，就会报警）
	deadlock.Opts.DeadlockTimeout = time.Second * 2

	// 2. 将普通的 sync.Mutex 替换为 deadlock.Mutex
	var muA, muB deadlock.Mutex

	// Goroutine 1
	go func() {
		muA.Lock()
		time.Sleep(100 * time.Millisecond)
		muB.Lock() // 这里会卡住
		muB.Unlock()
		muA.Unlock()
	}()

	// Goroutine 2（这里直接写在主协程，效果一样）
	muB.Lock()
	time.Sleep(100 * time.Millisecond)
	muA.Lock() // 这里也会卡住

	muA.Unlock()
	muB.Unlock()

	fmt.Println("程序结束")
}
