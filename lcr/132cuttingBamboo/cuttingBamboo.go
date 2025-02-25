package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(cuttingBamboo(12))
	fmt.Println(cuttingBamboo(120))
	fmt.Println(cuttingBamboo(5))
}

// 这种写法只在数据量小的时候可用，也就是不用取模时可用
func cuttingBamboo2(n int) int {
	// dp[i] 表示 i 的结果值
	// dp[i]: 正整数i对应的最大乘积
	dp := make(map[int]int)
	dp[0], dp[1], dp[2] = 0, 0, 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], dp[j]*(i-j), j*(i-j))
		}
	}
	return dp[n]
}

func cuttingBamboo(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	mod := int(math.Pow10(9)) + 7
	if b == 0 {
		// 分成b段,每段的长度是3
		return powN(3, a, mod)
	}
	if b == 1 {
		// 分成 a-1段长度是3，剩下的一段是4
		return powN(3, a-1, mod) * 4 % mod
	}
	// 分成 a段长度是3，剩下的一段是2
	return powN(3, a, mod) * 2 % mod
}

func powN(a, n, mod int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a % mod
	}
	c := n / 2
	d := powN(a, c, mod)

	ans := d * d % mod

	if n%2 == 1 {
		ans = (ans * a) % mod
	}
	return ans
}

// 最优： 3。把竹子尽可能切为多个长度为3的片段，留下的最后一段竹子的长度可能为0,1,2三种情况。
// 次优： 2。若最后一段竹子长度为2；则保留，不再拆为1+1。
// 最差： 1。若最后一段竹子长度为1；则应把一份3+1替换为2+2，因为2×2>3×1。

func countNumbers(cnt int) []int {
	mx := int(math.Pow10(cnt)) - 1
	ans := make([]int, mx)
	for i := 1; i <= mx; i++ {
		ans[i-1] = i
	}
	return ans
}
