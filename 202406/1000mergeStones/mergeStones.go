package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mergeStones([]int{3, 2, 4, 1}, 2))
	fmt.Println(mergeStones([]int{16, 43, 87, 30, 4, 98, 12, 30, 47, 45, 32, 4, 64, 14, 24, 84, 86, 51, 11, 22, 4}, 2))
}

func mergeStones(stones []int, k int) int {
	n := len(stones)
	sum := make([]int, n+1)

	for i := range stones {
		sum[i+1] = sum[i] + stones[i]
	}
	if (n-1)%(k-1) != 0 {
		return -1
	}
	mem := make([][][]int, n+1)
	for i := range mem {
		mem[i] = make([][]int, n+1)
		for j := range mem[i] {
			mem[i][j] = make([]int, k+1)
			for m := range mem[i][j] {
				mem[i][j][m] = -1
			}
		}
	}

	var dfs func(i, j, p int) int

	dfs = func(i, j, p int) int {

		if p <= 0 {
			return 0
		}

		if mem[i][j][p] != -1 {
			return mem[i][j][p]
		}
		if p == 1 {
			if i == j {

				mem[i][j][p] = 0
				return 0
			}
			a := dfs(i, j, k) + (sum[j+1] - sum[i])

			mem[i][j][p] = a
			return a
		}
		res := math.MaxInt
		for m := i; m < j; m = m + k - 1 {
			res = min(res, dfs(i, m, 1)+dfs(m+1, j, p-1))
		}
		mem[i][j][p] = res
		return res

	}
	return dfs(0, n-1, 1)
}
