package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(distinctSubseqII("abc"))
	fmt.Println(distinctSubseqII("z"))
}

func distinctSubseqII2(s string) int {
	n := len(s)
	mod := int(math.Pow10(9)) + 7
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 26)
	}
	dp[0][int(s[0])-int('a')] = 1
	for i := 1; i < n; i++ {
		c := s[i]
		idx := int(c) - int('a')
		dp[i] = dp[i-1]
		sum := 0
		for _, ch := range dp[i-1] {
			sum += ch
		}
		dp[i][idx] = (1 + sum) % mod
	}
	sum := 0
	for _, ch := range dp[n-1] {
		sum = (sum + ch) % mod
	}
	return sum
}

func distinctSubseqII(s string) int {
	mod := int(math.Pow10(9)) + 7
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 26)
	}
	// 初值
	dp[0][int(s[0])-int('a')] = 1
	for i := 1; i < n; i++ {
		idx := int(s[i]) - int('a')
		sum := 0
		for _, ch := range dp[i-1] {
			sum += ch
		}

		dp[i] = dp[i-1]
		dp[i][idx] = (1 + sum) % mod
	}
	sum := 0
	for _, ch := range dp[n-1] {
		sum = (sum + ch) % mod
	}
	return sum
}
