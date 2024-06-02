package main

import (
	"fmt"
)

func main() {
	fmt.Println(getLengthOfOptimalCompression("aaabcccd", 2))
	fmt.Println(getLengthOfOptimalCompression("aabbaa", 2))
	fmt.Println(getLengthOfOptimalCompression("aaaaaaaaaaa", 0))
}

func getLengthOfOptimalCompression(s string, k int) int {
	n := len(s)
	// 状态的定义为前i个字符（包括第 i 个字符）删掉j个之后的最短缩写长度
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = 2 * n // 求值最小，所以初始时定义一个较大值
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			if j > 0 {
				if i > 0 {
					dp[i][j] = min(dp[i][j], dp[i-1][j-1]) // 如果删除s[i]，那dp[i][j]就等于dp[i-1][j-1]
				} else {
					dp[i][j] = 0 // i==0，说明只有一个字符，此时如果 j>0那第就可以把这个字符删除了，最小值就是0
				}
			}
			// 如果保留s[i]，此时s[i]有可能参与压缩
			same := 0
			// 这个循环的意思就是从 i 开始，向前删除和 i 不相同的字符，求最小值
			for ii := i; ii >= 0; ii-- {
				if s[ii] == s[i] {
					same++
					cost := cst(same)
					if ii > 0 {
						cost += dp[ii-1][j-(i-ii+1-same)]
					}
					dp[i][j] = min(dp[i][j], cost)
				}
				if i-ii+1-same > j {
					break
				}
			}
		}
	}
	return dp[n-1][k]
}

func cst(x int) int {
	if x <= 1 {
		return x
	}
	return len(fmt.Sprintf("%d", x)) + 1
}
