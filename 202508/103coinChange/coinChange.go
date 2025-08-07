package main

import (
	"fmt"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

func coinChange2(coins []int, amount int) int {
	n := len(coins)

	var dfs func(i int, target int) int
	inf := 1 << 30
	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, amount+5)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i int, target int) int {
		if target < 0 {
			return 0
		}
		if i < 0 {
			if target == 0 {
				return 0
			}
			return inf
		}
		if mem[i][target] != -1 {
			return mem[i][target]
		}

		b := dfs(i-1, target)
		if target >= coins[i] {
			b = min(b, dfs(i, target-coins[i])+1)
		}
		mem[i][target] = b
		return b
	}
	ans := dfs(n-1, amount)
	if ans >= inf {
		return -1
	}
	return ans
}

func coinChange(coins []int, amount int) int {
	n := len(coins)
	f := make([][]int, n+5)
	inf := 1 << 30
	for i := range f {
		f[i] = make([]int, amount+5)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	f[0][0] = 0 // 初值
	for i := 0; i < n; i++ {
		for j := 0; j <= amount; j++ {
			f[i+1][j] = f[i][j]
			if j >= coins[i] {
				f[i+1][j] = min(f[i+1][j], f[i+1][j-coins[i]]+1)
			}
		}
	}
	if f[n][amount] >= inf {
		return -1
	}
	return f[n][amount]
}
