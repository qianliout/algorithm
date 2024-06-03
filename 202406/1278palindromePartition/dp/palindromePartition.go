package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(palindromePartition("abc", 2))
	fmt.Println(palindromePartition("aabbc", 3))
	fmt.Println(palindromePartition("leetcode", 8))
	fmt.Println(palindromePartition("fyhowoxzyrincxivwarjuwxrwealesxsimsepjdqsstfggjnjhilvrwwytbgsqbpnwjaojfnmiqiqnyzijfmvekgakefjaxryyml", 2))
}

/*
给你一个由小写字母组成的字符串 s，和一个整数 k。
请你按下面的要求分割字符串：

	首先，你可以将 s 中的部分字符修改为其他的小写英文字母。
	接着，你需要把 s 分割成 k 个非空且不相交的子串，并且每个子串都是回文串。

请返回以这种方式分割字符串所需修改的最少字符数。
*/
func palindromePartition(s string, k int) int {
	inf := math.MaxInt / 2
	n := len(s)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		// 初值
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if j == 1 {
				dp[i][j] = change(s, 0, i-1)
			} else {
				for m := j - 1; m < i; m++ {
					dp[i][j] = min(dp[i][j], dp[m][j-1]+change(s, m, i-1))
				}
			}
		}
	}
	if dp[n][k] >= inf {
		return 0
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
