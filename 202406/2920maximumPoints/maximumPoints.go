package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumPoints([][]int{{0, 1}, {1, 2}, {2, 3}}, []int{10, 10, 3, 3}, 5))
	fmt.Println(maximumPoints([][]int{{1, 0}, {2, 1}, {3, 1}}, []int{8, 2, 7, 1}, 2))
}

func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)
	g := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, 15)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	var dfs func(x, j, fa int) int

	dfs = func(x, j, fa int) int {
		if mem[x][j] >= 0 {
			return mem[x][j]
		}
		res1 := (coins[x] >> j) - k
		res2 := coins[x] >> (j + 1)
		for _, nx := range g[x] {
			if nx == fa {
				continue
			}
			res1 += dfs(nx, j, x)
			if j+1 < 14 { // 大于14之后结果就会是0，再 dfs 就没有意义
				res2 += dfs(nx, j+1, x)
			}
		}
		res := max(res1, res2)
		mem[x][j] = res
		return res
	}

	return dfs(0, 0, -1)
}
