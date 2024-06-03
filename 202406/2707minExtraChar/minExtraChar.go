package main

import (
	"fmt"
)

func main() {
	fmt.Println(minExtraChar("leetscode", []string{"leet", "code", "leetcode"}))
	fmt.Println(minExtraChar2("leetscode", []string{"leet", "code", "leetcode"}))
}

// 不加缓存会超时
func minExtraChar2(s string, dictionary []string) int {
	n := len(s)
	ss := []byte(s)
	dic := make(map[string]bool)
	for _, ch := range dictionary {
		dic[ch] = true
	}
	var dfs func(i int) int
	mem := make([]int, n+1)
	for i := range mem {
		mem[i] = -n
	}
	// i 表示s的前 i 个字符（包括 i）
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if mem[i] >= 0 {
			return mem[i]
		}
		// 不选 i，那么 i 就一定剩下下，所以在加1
		res := dfs(i-1) + 1
		for j := 0; j <= i; j++ {
			if dic[string(ss[j:i+1])] {
				res = min(res, dfs(j-1))
			}
		}
		mem[i] = res
		return res
	}
	return dfs(n - 1)
}

func minExtraChar(s string, dictionary []string) int {
	n := len(s)
	ss := []byte(s)
	// dp[i]表示 s[0:1)的结果，不包括i
	dp := make([]int, n+1)
	dp[0] = 0

	dic := make(map[string]bool)
	for _, ch := range dictionary {
		dic[ch] = true
	}
	for i := 1; i <= n; i++ {
		// 第i个字符不选
		dp[i] = dp[i-1] + 1
		// 贪心的思想，从0开始选，选最长的
		for j := 1; j <= i; j++ {
			if dic[string(ss[j:i])] {
				dp[i] = min(dp[i], dp[j-1])
			}
		}
	}
	return dp[n]
}
