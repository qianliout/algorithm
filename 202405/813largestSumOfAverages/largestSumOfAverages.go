package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestSumOfAverages([]int{9, 1, 2, 3, 9}, 3))
	fmt.Println(largestSumOfAverages([]int{4, 1, 7, 5, 6, 2, 3}, 4))
	fmt.Println(largestSumOfAverages([]int{1, 2, 3, 4, 5, 6, 7}, 4))
}

// 将 n 个元素划分为「最多」m 个连续段，最大化连续段的平均值之和。
func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	pre := make([]int, len(nums)+1)
	for i, ch := range nums {
		pre[i+1] = pre[i] + ch
	}
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, k+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, k); j++ {
			/*
				当 j=1 时，dp[i][j] 是对应区间 [0,i−1] 的平均值；
				当 j>1 时，我们将可以将区间 [0,i−1]分成 [0,x−1]] 和 [x,i−1] 两个部分，其中 x≥j−1，那么 dp[i][j] 等于所有这些合法的切分方式的平均值和的最大值。
			*/

			if j == 1 {
				dp[i][j] = float64(pre[i]) / float64(i)
				continue
			}
			for m := j - 1; m < i; m++ {
				dp[i][j] = max(dp[i][j], dp[m][j-1]+float64(pre[i]-pre[m])/float64(i-m))
			}
		}
	}
	return dp[n][k]
}
