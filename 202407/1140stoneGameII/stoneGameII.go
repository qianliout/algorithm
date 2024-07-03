package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(stoneGameII([]int{2, 7, 9, 4, 4}))
}

func stoneGameII(piles []int) int {
	n := len(piles)
	for i := n - 2; i >= 0; i-- {
		piles[i] += piles[i+1]
	}
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, n*2+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	var dfs func(i int, m int) int

	dfs = func(i, m int) int {
		if i >= n {
			return 0
		}
		if i+m*2 >= n { // 可以拿完时就全拿
			return piles[i]
		}
		if mem[i][m] >= 0 {
			return mem[i][m]
		}
		// 第一次拿最多，第二个人拿最少，这样就能保证第一个人最多了
		nn := math.MaxInt / 10
		for x := 1; x <= 2*m && i+x < n; x++ {
			nn = min(nn, dfs(i+x, max(m, x)))
		}
		mem[i][m] = piles[i] - nn
		return piles[i] - nn
	}

	return dfs(0, 1)
}
