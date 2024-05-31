package main

import (
	"math"
)

func main() {

}

func numTilings(n int) int {
	dp := make([]int, n+10) // 这里最好多建几个，因为 n 可能小于3，下面又赋值给3，所以可能会出错
	base := int(math.Pow(10, 9)) + 7
	dp[0], dp[1], dp[2], dp[3] = 0, 1, 2, 5
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1]*2 + dp[i-3]
		dp[i] = dp[i] % base
	}

	return dp[n]
}
