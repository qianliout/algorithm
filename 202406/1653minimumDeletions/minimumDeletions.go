package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(minimumDeletions("aababbab"))
	fmt.Println(minimumDeletions("aababbab"))
}

func minimumDeletions1(s string) int {
	bc := 0 // 到 i 时 字符 b 的个数
	// 表示使 s 的前 i 个字母平衡的最少删除次数：
	dp := 0 // dp数组，因为dp[i]只和 dp[i-1]相关，所以可以只用一个变量
	for _, ch := range s {
		if ch == 'b' {
			// 到i 时最后一个字符是 b，此时不用删除dp[i] = dp[i-1]
			bc++
		} else {
			// 1，保留 a 那么得把之前的字符 b 全删除
			// 2,删除这个 a,那么 dp[i] = dp[i-1]+1
			dp = min(dp+1, bc)
		}
	}
	return dp
}

func minimumDeletions(s string) int {
	/*
		因此，我们维护两个变量 lb 和 ra 分别表示 s[0,..,i−1] 中字符 bbb 的个数以及 s[i+1,..n−1]中字符 a 的个数，那么我们需要删除的字符数为 lb+ra。枚举过程中，更新变量 lb 和 ra。
	*/
	lb := 0
	ra := strings.Count(s, "a")

	ans := len(s)
	for _, ch := range s {
		if ch == 'a' {
			ra--
		}
		ans = min(ans, ra+lb)
		if ch == 'b' {
			lb++
		}
	}
	return ans
}

func minimumDeletions2(s string) int {
	ac := strings.Count(s, "a")
	ans, del := ac, ac

	for _, ch := range s {
		if ch == 'a' {
			del--
		} else if ch == 'b' {
			del++
		}
		ans = min(ans, del)
	}
	return ans
}
