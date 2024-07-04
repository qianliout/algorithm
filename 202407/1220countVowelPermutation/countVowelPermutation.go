package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countVowelPermutation(5))
}

func countVowelPermutation1(n int) int {
	if n <= 0 {
		return 1
	}

	mod := int(math.Pow10(9)) + 7
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 5)
	}
	for i := 0; i < 5; i++ {
		dp[0][i] = 1
	}

	// aeiou
	for i := 0; i < n-1; i++ {
		// 每个元音 'a' 后面都只能跟着 'e'
		dp[i+1][1] = (dp[i+1][1] + dp[i][0]) % mod
		// 每个元音 'e' 后面只能跟着 'a' 或者是 'i'
		dp[i+1][0] = (dp[i+1][0] + dp[i][1]) % mod
		dp[i+1][2] = (dp[i+1][2] + dp[i][1]) % mod
		// 每个元音 'i' 后面 不能 再跟着另一个 'i'
		dp[i+1][0] = (dp[i+1][0] + dp[i][2]) % mod
		dp[i+1][1] = (dp[i+1][1] + dp[i][2]) % mod
		dp[i+1][3] = (dp[i+1][3] + dp[i][2]) % mod
		dp[i+1][4] = (dp[i+1][4] + dp[i][2]) % mod
		// 每个元音 'o' 后面只能跟着 'i' 或者是 'u'
		dp[i+1][2] = (dp[i+1][2] + dp[i][3]) % mod
		dp[i+1][4] = (dp[i+1][4] + dp[i][3]) % mod
		// 每个元音 'u' 后面只能跟着 'a'
		dp[i+1][0] = (dp[i+1][0] + dp[i][4]) % mod
	}
	ans := 0
	for i := 0; i < 5; i++ {
		ans = (ans + dp[n-1][i]) % mod
	}
	return ans % mod
}

func countVowelPermutation(n int) int {
	if n <= 0 {
		return 1
	}

	mod := int(math.Pow10(9)) + 7
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 5)
	}
	for i := 0; i < 5; i++ {
		dp[0][i] = 1
	}

	// aeiou
	for i := 0; i < n-1; i++ {
		// 每个元音 'a' 后面都只能跟着 'e'
		dp[i+1][1] += dp[i][0]
		// 每个元音 'e' 后面只能跟着 'a' 或者是 'i'
		dp[i+1][0] += dp[i][1]
		dp[i+1][2] += dp[i][1]
		// 每个元音 'i' 后面 不能 再跟着另一个 'i'
		dp[i+1][0] += dp[i][2]
		dp[i+1][1] += dp[i][2]
		dp[i+1][3] += dp[i][2]
		dp[i+1][4] += dp[i][2]
		// 每个元音 'o' 后面只能跟着 'i' 或者是 'u'
		dp[i+1][2] += dp[i][3]
		dp[i+1][4] += dp[i][3]
		// 每个元音 'u' 后面只能跟着 'a'
		dp[i+1][0] += dp[i][4]
		for j := 0; j < 5; j++ {
			dp[i+1][j] %= mod
		}
	}
	ans := 0
	for i := 0; i < 5; i++ {
		ans = (ans + dp[n-1][i]) % mod
	}
	return ans % mod
}
