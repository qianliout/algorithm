package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minCost([]int{1, 2, 1, 2, 1, 3, 3}, 2))
	fmt.Println(minCost([]int{1, 2, 1, 2, 1}, 2))
	fmt.Println(minCost([]int{68, 81, 76, 91, 4, 0, 50, 5, 66, 23, 27, 91, 37, 61, 2, 0, 78, 11, 76, 58, 47, 46, 91, 57, 12, 61, 12, 17, 64, 26, 92, 71, 32, 52, 6, 70, 44, 77, 50, 7, 50, 2, 77, 56, 83, 69, 41, 88, 18, 2, 80, 2, 1, 2, 1, 88, 76, 88, 3, 64, 75, 93, 5, 18, 11, 65, 75, 16, 75, 17, 91, 2, 57, 10, 3, 29, 4, 83, 75, 15, 21, 20, 93, 32, 9, 29, 71, 21, 0, 93, 66, 52, 24, 90, 53, 82, 62, 85, 52, 27, 41, 58, 30, 3, 79, 63, 41, 50, 8, 71, 61, 59, 94, 13, 32, 61, 52, 80, 55, 84, 94, 61, 4, 19, 36, 25, 77, 43, 80, 72, 38, 17, 76, 74, 2, 31, 58, 39, 55, 39, 35, 88, 54, 53, 35, 80, 45, 8, 61, 67, 3, 48, 66, 15, 9, 22, 81, 30, 15, 48, 54, 87, 59, 9, 54, 65, 74, 13, 54, 48, 78, 55, 48, 93, 88, 94, 58, 84, 36, 44, 19, 50, 71, 54, 17, 76, 27, 58, 78, 88}, 178))
}

// 找出并返回拆分 nums 的所有可行方案中的最小代价。

func minCost(nums []int, k int) int {
	inf := math.MaxInt / 2
	n := len(nums)
	// dp[i][j]表示 nums[:i]分隔成 j 次个子数组，注意这里不包括 i 这个字符
	dp := make([][]int, n+1)

	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	for i := 0; i <= n; i++ {
		for j := 1; j <= min(i, n); j++ {
			if j == 1 {
				dp[i][j] = cal(nums, 0, i-1, k, mem)
				continue
			}
			for m := j - 1; m < i; m++ {
				dp[i][j] = min(dp[i][j], dp[m][j-1]+cal(nums, m, i-1, k, mem))
			}
		}
	}
	ans := inf
	for _, ch := range dp[n] {
		ans = min(ans, ch)
	}
	return ans
}

func cal(nums []int, le, ri int, k int, mem [][]int) int {
	if mem[le][ri] >= 0 {
		return mem[le][ri]
	}
	exit := make(map[int]int)
	for i := le; i <= ri; i++ {
		exit[nums[i]]++
	}
	ans := k
	for _, v := range exit {
		if v >= 2 {
			ans += v
		}
	}
	mem[le][ri] = ans
	return ans
}
