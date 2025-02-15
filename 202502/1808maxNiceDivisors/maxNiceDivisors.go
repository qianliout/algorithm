package main

func main() {

}

func maxNiceDivisors(primeFactors int) int {
	n := primeFactors
	if n <= 3 {
		return n
	}
	mod := int(1e9 + 7)
	ans := 0
	q, r := n/3, n%3
	if r == 0 {
		ans = powN(3, q, mod)
	} else if r == 1 {
		ans = powN(3, q-1, mod) * 4
	} else {
		ans = powN(3, q, mod) * 2
	}
	return ans % mod
}

// 没有能理解
// https://leetcode.cn/problems/maximize-number-of-nice-divisors/solutions/684836/fan-yi-wan-zhi-hou-jiu-xiang-dang-yu-yua-113z/
func powN(a, n int, mod int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a % mod
	}
	b := n / 2
	c := powN(a, b, mod)
	ans := c * c % mod
	if n%2 == 1 {
		ans = ans * a % mod
	}
	return ans
}

func integerBreak(n int) int {
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
