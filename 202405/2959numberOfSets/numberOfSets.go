package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numberOfSets(3, 5, [][]int{{0, 1, 2}, {1, 2, 10}, {0, 2, 10}}))
}

func numberOfSets(n int, maxDistance int, roads [][]int) int {
	inf := math.MaxInt / 2
	// g := make([][]int, n)
	// for i := range g {
	// 	g[i] = make([]int, n)
	// 	for j := range g[i] {
	// 		if i != j {
	// 			g[i][j] = inf
	// 		}
	// 	}
	// }
	// for _, ch := range roads {
	// 	x, y, z := ch[0], ch[1], ch[2]
	// 	g[x][y] = min(g[x][y], z)
	// 	g[y][x] = min(g[y][x], z)
	// }

	ans := 0
	for mask := 0; mask < 1<<n; mask++ {
		a := check(roads, n, mask, maxDistance, inf)
		if a == 1 {
			fmt.Println("check", mask)
		}
		ans += a
		// ans += check(roads, n, mask, maxDistance, inf)
	}

	return ans
}

func check(roads [][]int, n int, mask int, maxDistance int, inf int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			if i != j {
				dp[i][j] = inf
			}
		}
	}
	for i := 0; i < n; i++ {
		dp[i][i] = 0
	}

	for _, ch := range roads {
		x, y, z := ch[0], ch[1], ch[2]
		dp[x][y] = min(dp[x][y], z)
		dp[y][x] = min(dp[y][x], z)
	}

	// dp[0] = g
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// # 注意x,y,z都保留的时候，才能用floyd的公式更新最短距离
				if (mask>>k)&1 == 1 && (mask>>i)&1 == 1 && (mask>>j)&1 == 1 {
					dp[i][j] = min(dp[i][j], dp[i][k]+dp[k][j])
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// if里的内容成立，就说明存在2个保留的分部距离超标，为非法方案
			if (mask>>i)&1 == 1 && (mask>>j)&1 == 1 && dp[i][j] > maxDistance {
				return 0
			}
		}
	}
	return 1
}
