package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(distributeCookies([]int{8, 15, 10, 20, 8}, 2))
	fmt.Println(distributeCookies([]int{1, 1, 1}, 2))
}

func distributeCookies2(cookies []int, k int) int {
	n := len(cookies)
	m := 1 << n

	sum := make([]int, m)
	// i 是以二进制表示的集合，sum 就表示的各个不同的集合的值
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				sum[i] += cookies[j]
			}
		}
	}
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	inf := math.MaxInt / 10
	// 初值
	// k 的下标从0开始，

	dp[1] = sum // 没有能理解
	for i := 1; i < k; i++ {
		for j := 0; j < m; j++ {
			dp[i][j] = inf
			for s := j; s > 0; s = (s - 1) & j {
				dp[i][j] = min(dp[i][j], max(dp[i-1][j^s], sum[s]))
			}
		}
	}
	return dp[k-1][m-1]
}

func distributeCookies(cookies []int, k int) int {
	n := len(cookies)
	m := 1 << n

	sum := make([]int, m)
	// i 是以二进制表示的集合，sum 就表示的各个不同的集合的值
	// 例如 cookies= [1,2,3]
	// 那么2的二进制是 010 表示选中间的2 值值，011表示 选2，3
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				sum[i] += cookies[j]
			}
		}
	}
	// dp[i]表示分成i组的值
	// dp[i][j]表示集合 j 的所有元素，分成 i 组，的值
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	inf := math.MaxInt / 10
	// dp[0] 没有新意义，当然有些题解是从下标0开始算到 k-1,我这里从下标1始算到 k,本质是一样的
	// 初值
	dp[1] = sum // 没有能理解
	for i := 2; i <= k; i++ {
		for j := 0; j < m; j++ {
			dp[i][j] = inf // 后期取最小值所以这里初值是一个最大值
			// s 是 j 子集，当然包括 j 本身
			for s := j; s > 0; s = (s - 1) & j {
				dp[i][j] = min(dp[i][j], max(dp[i-1][j^s], sum[s]))
			}
		}
	}
	return dp[k][m-1]
}
