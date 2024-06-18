package main

import (
	"math"
)

func main() {

}

func numPermsDISequence(s string) int {
	// 令dp[i][j]表示字符串长度为i时，以j为结尾的序列个数
	mod := int(math.Pow10(9)) + 7
	n := len(s)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] == 'D' {
			dp[i][i] = 0 // 最后一个数是i，肯定不可能构成降序
			for j := i - 1; j >= 0; j-- {
				dp[i][j] = (dp[i][j+1] + dp[i-1][j]) % mod
			}
		} else {
			dp[i][0] = 0 // 最后一个数是0，肯定不可能构成升序
			for j := 1; j <= i; j++ {
				dp[i][j] = (dp[i-1][j-1] + dp[i][j-1]) % mod
			}
		}
	}
	ans := 0
	for j := 0; j <= n; j++ {
		ans = (ans + dp[n][j]) % mod
	}
	return ans
}
