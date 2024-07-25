package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countGoodNumbers(50)) // 564908303
	fmt.Println(countGoodNumbers(4))  // 564908303
}

// 直接这样写有问题，需要快速幂
func countGoodNumbers2(n int64) int {
	dp := make([]int, n)
	dp[0] = 5 // 可以有前导0
	mod := int(math.Pow10(9)) + 7

	for i := 1; i < int(n); i++ {
		if i%2 == 1 {
			dp[i] = dp[i-1] * 4
		} else {
			dp[i] = dp[i-1] * 5
		}
		dp[i] = dp[i] % mod
	}

	return dp[n-1] % mod
}

// 这样写不内存，但是超时
func countGoodNumbers3(n int64) int {
	pre := 5 // 可以有前导0
	mod := int(math.Pow10(9)) + 7

	for i := 1; i < int(n); i++ {
		if i%2 == 1 {
			pre = pre * 4
		} else {
			pre = pre * 5
		}
		pre %= mod
	}

	return pre % mod
}

// 一开始为5
// 每次到奇数就 * 5
// 每次到偶数就 * 4
// 所以结果为
// 4 ^ (n中的偶数) * 5 ^ (n中的偶数)

func countGoodNumbers(n int64) int {
	a := (n + 1) / 2
	b := n / 2
	mod := int(math.Pow10(9)) + 7
	c := pow(int(a), 5, mod) * pow(int(b), 4, mod)
	return c % mod
}

// 快速幂
func pow(n, x, mod int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	nex := pow(n/2, x, mod)
	a := nex * nex
	if n&1 == 1 {
		a = a * x
	}
	return a % mod
}
