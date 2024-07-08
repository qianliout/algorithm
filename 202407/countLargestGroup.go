package main

import (
	"fmt"
)

func main() {
	fmt.Println(countLargestGroup(13))
}

func countLargestGroup(n int) int {
	rec := make(map[int]int)
	maxCount := 0
	for i := 1; i <= n; i++ {
		key := help(i)
		rec[key]++

	}
	for _, count := range rec {
		if count > maxCount {
			maxCount = count
		}
	}

	// 统计最大计数出现的次数
	actualMaxCount := 0
	for _, count := range rec {
		if count == maxCount {
			actualMaxCount++
		}
	}

	return actualMaxCount
}

func help(num int) int {
	ans := 0
	for num > 0 {
		ans += ans % 10
		num = num / 10
	}
	return ans
}
