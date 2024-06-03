package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(palindromePartition("abc", 2))
	fmt.Println(palindromePartition("aabbc", 3))
	fmt.Println(palindromePartition("leetcode", 8))
	fmt.Println(palindromePartition("fyhowoxzyrincxivwarjuwxrwealesxsimsepjdqsstfggjnjhilvrwwytbgsqbpnwjaojfnmiqiqnyzijfmvekgakefjaxryyml", 32))
}

func palindromePartition(s string, k int) int {
	inf := math.MaxInt / 2
	n := len(s)
	// dp[i][j]表示 s[:i]分隔成 j 个字符串的结果集，注意这里不包括 i 这个字符
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if j == 1 {
				dp[i][j] = change(s, 0, i-1)
				continue
			}
			// 前面还需要分隔成j-1段，所以至少需要j-1个字符
			// 所以 s[0:j-1] 要留给前面的 j-1段，从 j-1开始
			// 易错点2： m<i,而不能是 m<=i,因为题中说了，不能有空串，如果 m=i，那么就是一个空串
			for m := j - 1; m < i; m++ {
				dp[i][j] = min(dp[i][j], dp[m][j-1]+change(s, m, i-1))
			}
		}
	}
	return dp[n][k]
}

func change(ss string, le, ri int) int {
	ans := 0
	for le < ri {
		if ss[le] != ss[ri] {
			ans++
		}
		le++
		ri--
	}
	return ans
}
