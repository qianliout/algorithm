package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	runtime.GC()                                 // 手动触发 GC（仅用于测试）
	fmt.Println("Before a[:0]:", len(a), cap(a)) // 5, 5

	a = a[:0]
	runtime.GC()
	fmt.Println("After a[:0]:", len(a), cap(a)) // 0, 5（底层数组仍在）

	a = nil
	runtime.GC()
	fmt.Println("After nil:", len(a), cap(a)) // 0, 0（底层数组可被回收）
}
