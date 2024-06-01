package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numberOfSets(3, 5, [][]int{{0, 1, 2}, {1, 2, 10}, {0, 2, 10}}))
}

// 有错，但是不知道是什么原因
func numberOfSets(n int, maxDistance int, roads [][]int) int {
	inf := math.MaxInt / 2
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			if i != j {
				g[i][j] = inf
			}
		}
	}
	for _, ch := range roads {
		x, y, z := ch[0], ch[1], ch[2]
		g[x][y] = min(g[x][y], z)
		g[y][x] = min(g[y][x], z)
	}

	ans := 0
	for mask := 0; mask < 1<<n; mask++ {
		a := check(g, n, mask, maxDistance, inf)
		if a == 1 {
			fmt.Println("check", mask)
		}
		// ans += a
		ans += check(g, n, mask, maxDistance, inf)
	}

	return ans
}

func check(g [][]int, n int, mask int, maxDistance int, inf int) int {
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}
	dp[0] = g
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// # 注意x,y,z都保留的时候，才能用floyd的公式更新最短距离
				if (mask>>k)&1 == 1 && (mask>>i)&1 == 1 && (mask>>j)&1 == 1 {
					dp[k+1][i][j] = min(dp[k][i][j], dp[k][i][k]+dp[k][k][j])
				}
			}
		}
	}
	for k := 0; k <= n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// if里的内容成立，就说明存在2个保留的分部距离超标，为非法方案
				if (mask>>i)&1 == 1 && (mask>>j)&1 == 1 && dp[k][i][j] < inf && dp[k][i][j] > maxDistance {
					return 0
				}
			}
		}
	}
	return 1
}
