package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(coinChange([]int{156, 265, 40, 280}, 9109))
}

func coinChange2(coins []int, amount int) int {
	cnt := make(map[int]bool)
	mx := 0
	for _, ch := range coins {
		mx = max(mx, ch)
		cnt[ch] = true
	}
	dp := make([]int, amount+1)

	// 这样写法是对的，但是会超时
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for j := i - 1; j >= 0 && i-j <= mx; j-- {
			if cnt[i-j] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	cnt := make(map[int]bool)
	mx := 0
	for _, ch := range coins {
		mx = max(mx, ch)
		cnt[ch] = true
	}
	dp := make([]int, amount+1)

	// 这样写法是对的，但是会超时
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, ch := range coins {
			// 因为上面排序了
			if i-ch < 0 {
				break
			}
			dp[i] = min(dp[i], dp[i-ch]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
