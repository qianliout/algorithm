package main

import (
	"fmt"
)

func main() {
	fmt.Println(soupServings(850))
}

func soupServings(n int) float64 {
	n = (n + 24) / 25
	if n >= 179 {
		return 1
	}
	mem := make([][]float64, n+10)
	for i := range mem {
		mem[i] = make([]float64, n+10)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	return dfs(n, n, mem)
}

func dfs(a, b int, mem [][]float64) float64 {

	if a <= 0 && b <= 0 {
		return 0.5
	}
	if a <= 0 && b > 0 {
		return 1
	}
	// 此时不管怎么样都是B 先分配完
	if b <= 0 {
		return 0
	}
	if mem[a][b] >= 0 {
		return mem[a][b]
	}
	res := (dfs(a-4, b, mem) + dfs(a-3, b-1, mem) +
		dfs(a-2, b-2, mem) + dfs(a-1, b-3, mem)) / 4
	mem[a][b] = res
	return res
}
