package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println(dieSimulator(2, []int{1, 1, 1, 1, 1, 1}))
	fmt.Println(dieSimulator(2, []int{1, 1, 2, 2, 2, 3}))
	fmt.Println(dieSimulator(3, []int{1, 1, 1, 2, 2, 3}))
}

func dieSimulator(n int, rollMax []int) int {
	var dfs func(n, i, j int) int
	// k 表示还剩余多少次机会
	// i表示前一个数字是多少
	// j 表示前一个数字已连续了多少次
	mod := int(math.Pow10(9)) + 7
	mc := slices.Max(rollMax)

	mem := make([][][]int, n+1)
	for i := range mem {
		mem[i] = make([][]int, 7)
		for j := range mem[i] {
			mem[i][j] = make([]int, mc+1)
			for k := range mem[i][j] {
				mem[i][j][k] = -1
			}
		}
	}

	dfs = func(k, i, j int) int {
		if k <= 0 {
			return 1
		}
		if mem[k][i][j] != -1 {
			return mem[k][i][j]
		}
		res := 0
		for num := 0; num < 6; num++ {
			if num != i {
				res += dfs(k-1, num, 1)
			} else {
				if rollMax[i] > j {
					// 如果没有
					res += dfs(k-1, num, j+1)
				}
			}
		}
		mem[k][i][j] = res % mod
		return res % mod
	}
	res := 0
	for i := 0; i < 6; i++ {
		res += dfs(n-1, i, 1)
		res %= mod
	}
	return res % mod
}
