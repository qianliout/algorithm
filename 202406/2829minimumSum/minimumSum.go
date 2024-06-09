package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumSum(5, 4))
	fmt.Println(minimumSum(2, 6))
}

func minimumSum2(n int, k int) int {
	m := k / 2
	ans := 0
	for i := 1; i <= min(n, m); i++ {
		ans += i
	}
	for i := k; i <= k+(n-m)-1; i++ {
		ans += i
	}

	return ans
}

func minimumSum(n int, k int) int {
	m := min(n, k/2)
	ans := 0
	ans += (1 + m) * m
	ans += (k + k + n - m - 1) * (n - m)
	return ans / 2
}
