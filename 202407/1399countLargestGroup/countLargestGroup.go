package main

import (
	"fmt"
)

func main() {
	fmt.Println(countLargestGroup(13))
}

func countLargestGroup(n int) int {
	rec := make(map[int]int)
	for i := 1; i <= n; i++ {
		key := help(i)
		rec[key]++
	}

	maxCount := 0
	maxValue := -1
	cnt2 := make(map[int]int)
	for _, count := range rec {
		cnt2[count]++
	}

	// 统计最大计数出现的次数
	for k, count := range cnt2 {
		if k > maxCount {
			maxCount = k
			maxValue = count
		}
	}
	return maxValue
}

func help(num int) int {
	ans := 0
	for num > 0 {
		ans += num % 10
		num = num / 10
	}
	return ans
}
