package main

import "fmt"

func main() {
	fmt.Println(largestPalindrome(3, 5))
}

func largestPalindrome(n int, k int) string {
	pow := make([]int, n)
	pow[0] = 1
	// for i := 1; i < n; i++ {
	// 	pow[i] = pow[i-1] * 10 % k
	// }
	for i := 0; i < n; i++ {
		pow[i] = pow10(i, k)
	}

	fmt.Println(pow)

	ans := make([]byte, n)
	m := (n + 1) / 2 // 向上取整数

	visit := make([][]bool, m+1)
	for i := range visit {
		visit[i] = make([]bool, k)
	}
	// i 表示 前 i 位,j 表示模k后的结果
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i >= m {
			return j == 0
		}
		visit[i][j] = true
		for d := 9; d >= 0; d-- {
			j2 := j
			if n%2 == 1 && i == m-1 { // 正中间
				j2 = (j + d*pow[i]) % k
			} else {
				j2 = (j + d*pow[i] + d*pow[n-i-1]) % k
			}
			if !visit[i+1][j2] && dfs(i+1, j2) {
				ans[i] = byte('0' + d)
				ans[n-i-1] = ans[i]
				return true
			}
		}
		return false
	}
	dfs(0, 0)
	return string(ans)
}

// 快速幂
func pow10(x int, k int) int {
	if x == 0 {
		return 1
	}
	if x == 1 {
		return 10 % k
	}
	a := pow10(x>>1, k)
	a = a * a % k
	if x&1 == 1 {
		a = a * 10 % k
	}
	return a % k
}
