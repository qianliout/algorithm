package main

import (
	"math"
	"slices"
)

func main() {

}

func minimizeTheDifference(mat [][]int, target int) int {
	mx := 0
	for i := range mat {
		mx += slices.Min(mat[i])
	}
	mx += target

	n := len(mat)
	m := len(mat[0])
	inf := math.MaxInt64 / 100

	var dfs func(i, s int) int
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, mx+10)
	}

	dfs = func(i, j int) int {
		if i < 0 || i >= n {
			return abs(j - target)
		}
		if j > mx {
			return inf
		}
		if mem[i][j] != 0 {
			return mem[i][j]
		}

		ans := inf
		for k := 0; k < m; k++ {
			ans = min(ans, dfs(i+1, j+mat[i][k]))
		}
		mem[i][j] = ans
		return ans
	}
	return dfs(0, 0)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
