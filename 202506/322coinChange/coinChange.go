package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

func coinChange2(coins []int, amount int) int {
	return dfs(coins, len(coins)-1, amount)
}

func dfs(coins []int, i int, t int) int {
	if i < 0 {
		if t == 0 {
			return 0
		}
		return math.MaxInt32
	}
	c := coins[i]
	ans := dfs(coins, i-1, t)
	if t-c >= 0 {
		ans = min(ans, dfs(coins, i, t-c)+1)
	}
	return ans
}

func coinChange(coins []int, amount int) int {

	n := len(coins)
	inf := 1 << 32
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, amount+1)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	f[0][0] = 0
	for i := 0; i < n; i++ {
		for t := 0; t <= amount; t++ {
			c := coins[i]
			f[i+1][t] = f[i][t]
			if t-c >= 0 {
				f[i+1][t] = min(f[i+1][t], f[i+1][t-c]+1)
			}
		}
	}
	if f[n][amount] == inf {
		return -1
	}
	return f[n][amount]
}
