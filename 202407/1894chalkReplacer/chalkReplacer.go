package main

import (
	"fmt"
)

func main() {
	fmt.Println(chalkReplacer([]int{1}, 10000000))
}

// 直接写会超时
func chalkReplacer2(chalk []int, k int) int {
	i, n := 0, len(chalk)
	for {
		idx := i % n
		if k < chalk[idx] {
			return idx
		}
		k -= chalk[idx]
		i++
	}
}

func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for _, ch := range chalk {
		sum += ch
	}
	k = k % sum
	i, n := 0, len(chalk)
	for {
		idx := i % n
		if k < chalk[idx] {
			return idx
		}
		k -= chalk[idx]
		i++
	}
}
