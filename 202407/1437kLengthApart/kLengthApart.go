package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(kLengthApart([]int{1, 0, 0, 0, 1, 0, 0, 1}, 2))
	fmt.Println(kLengthApart([]int{1, 0, 0, 1, 0, 1}, 2))
}

func kLengthApart(nums []int, k int) bool {
	// 可以认为第一个1在无限远，这里除以10，是为防止越界
	first := math.MinInt / 10
	ans := math.MaxInt / 10
	for i, ch := range nums {
		if ch == 1 {
			ans = min(i-first-1, ans)
			first = i
		}
	}
	return ans >= k
}
