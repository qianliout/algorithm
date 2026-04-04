package main

import "fmt"

func main() {
	chPtr := new(chan int)

	fmt.Println("尝试从 nil channel 接收数据...")

	// 如果直接在主协程读取 nil channel，整个程序会立刻陷入死锁
	val := <-(*chPtr) // ⛔ 程序卡死，触发内置保护报错: fatal error: all goroutines are asleep - deadlock!

	fmt.Println("接收到的值:", val)
}
