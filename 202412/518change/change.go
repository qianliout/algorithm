package main

import (
	"fmt"
)

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
	fmt.Println(change1(5, []int{1, 2, 5}))
}

func change1(amount int, coins []int) int {
	var dfs func(x int, c int) int
	n := len(coins)
	dfs = func(i int, c int) int {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}
		if c < 0 {
			return 0
		}
		// if c < coins[i] {
		// 	return dfs(i-1, c)
		// }
		a := dfs(i, c-coins[i])
		b := dfs(i-1, c)
		return a + b
	}
	ans := dfs(n-1, amount)
	return ans
}

func change(amount int, coins []int) int {
	n := len(coins)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, amount+1)
	}
	for i := range f {
		f[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j >= coins[i-1] {
				f[i][j] = f[i-1][j] + f[i][j-coins[i-1]]
			} else {
				f[i][j] = f[i-1][j]
			}
		}
	}

	return f[n][amount]
}
