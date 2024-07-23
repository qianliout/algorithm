package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getMaximumConsecutive([]int{1, 1, 1, 4}))
}

// 请返回从 0 开始（包括 0 ），你最多能 构造 出多少个连续整数。
// 这个题目要求从0开始，这是这个题目的关键
func getMaximumConsecutive(coins []int) int {
	start := 0
	sort.Ints(coins)
	for _, c := range coins {
		if c > start+1 {
			break
		}
		start = start + c
	}

	return start + 1
}
