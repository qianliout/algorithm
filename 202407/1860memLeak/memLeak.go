package main

import (
	"fmt"
)

func main() {
	fmt.Println(memLeak(2, 2))
}

func memLeak(memory1 int, memory2 int) []int {
	start := 1
	ti := 1
	for memory1 >= start || memory2 >= start {
		if memory1 >= memory2 {
			memory1 -= start
			start++
			ti++
		} else {
			memory2 -= start
			start++
			ti++
		}
	}
	return []int{ti, memory1, memory2}
}
