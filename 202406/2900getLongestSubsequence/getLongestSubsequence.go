package main

import (
	"fmt"
)

func main() {
	fmt.Println(getWordsInLongestSubsequence([]string{"bab", "dab", "cab"}, []int{1, 2, 2}))
}

func getLongestSubsequence(words []string, groups []int) []string {
	ans := make([]string, 0)
	for i, x := range groups {
		if i == len(words)-1 || x != groups[i+1] {
			ans = append(ans, words[i])
		}
	}
	return ans
}

/*
定义 f[i]示从 i 到 n−1 中，我们选出的最长子序列的长度（第一个下标一定是 i）。定义成后缀是为了方便后面输出具体方案。

初始值 f[i]=1，表示选择它自己作为子序列。
*/
func getWordsInLongestSubsequence2(words []string, groups []int) []string {
	n := len(words)
	from := make([]int, n)
	dp := make([]int, n)
	// 初值
	// for i := range dp {
	// 	dp[i] = 1 // 以自已做为子序列
	// }
	mx := n - 1
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if dp[j] > dp[i] && groups[j] != groups[i] &&
				len(words[i]) == len(words[j]) && ham(words[i], words[j]) == 1 {
				dp[i] = dp[j]
				from[i] = j
			}
		}
		dp[i]++

		if dp[i] > dp[mx] {
			mx = i
		}
	}
	ans := make([]string, dp[mx])
	for i := range ans {
		ans[i] = words[mx]
		mx = from[mx]
	}
	return ans
}

func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(words)

	from := make([]int, n)
	dp := make([]int, n)
	// 初值
	for i := range dp {
		dp[i] = 1 // 以自已做为子序列
	}
	mx := n - 1
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if dp[j] >= dp[i] && groups[j] != groups[i] &&
				len(words[i]) == len(words[j]) && ham(words[i], words[j]) == 1 {
				dp[i] = max(dp[i], dp[j]+1)

				from[i] = j
			}
		}

		if dp[i] > dp[mx] {
			mx = i
		}
	}
	// 这种输出结果的方式要学习
	ans := make([]string, dp[mx])
	for i := range ans {
		ans[i] = words[mx]
		mx = from[mx]
	}
	return ans
}

func ham(s1, s2 string) int {
	if len(s1) > len(s2) {
		return ham(s2, s1)
	}
	ans := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			ans++
		}
	}
	ans += len(s2) - len(s1)
	return ans
}
