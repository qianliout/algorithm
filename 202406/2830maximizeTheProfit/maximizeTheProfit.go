package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximizeTheProfit(5, [][]int{{0, 0, 1}, {0, 2, 2}, {1, 3, 2}}))
	fmt.Println(maximizeTheProfit(5, [][]int{{0, 0, 1}, {0, 2, 10}, {1, 3, 2}}))
	fmt.Println(maximizeTheProfit(5, [][]int{{0, 2, 4}}))

}

// 使用记忆化dfs 都会超时
func maximizeTheProfit2(n int, offers [][]int) int {
	sort.Slice(offers, func(i, j int) bool {
		if offers[i][1] < offers[j][1] {
			return true
		} else if offers[i][1] > offers[j][1] {
			return false
		}
		return offers[i][0] < offers[j][0]
	})

	mem := make(map[int]int)

	var dfs func(i int) int

	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if i == 0 {
			return offers[0][2]
		}
		if va, ok := mem[i]; ok {
			return va
		}
		res1 := dfs(i - 1) // 不选
		// 选
		res2 := offers[i][2]
		for j := i - 1; j >= 0; j-- {
			if offers[j][1] < offers[i][0] {
				res2 += dfs(j)
				break
			}
		}
		res := max(res1, res2)
		mem[i] = res
		return res
	}
	return dfs(len(offers) - 1)
}

func maximizeTheProfit3(n int, offers [][]int) int {
	sort.Slice(offers, func(i, j int) bool {
		if offers[i][1] < offers[j][1] {
			return true
		} else if offers[i][1] > offers[j][1] {
			return false
		}
		return offers[i][0] < offers[j][0]
	})

	mem := make([]int, max(n, len(offers))+10)
	for i := range mem {
		mem[i] = -1
	}

	var dfs func(i int) int

	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if i == 0 {
			return offers[0][2]
		}
		if mem[i] >= 0 {
			return mem[i]
		}

		res1 := dfs(i - 1) // 不选
		// 选
		res2 := offers[i][2]
		for j := i - 1; j >= 0; j-- {
			if offers[j][1] < offers[i][0] {
				res2 += dfs(j)
				break
			}
		}
		res := max(res1, res2)
		mem[i] = res
		return res
	}
	return dfs(len(offers) - 1)
}

// 还是超时
func maximizeTheProfit4(n int, offers [][]int) int {
	if n == 0 || len(offers) == 0 {
		return 0
	}
	sort.Slice(offers, func(i, j int) bool {
		if offers[i][1] < offers[j][1] {
			return true
		} else if offers[i][1] > offers[j][1] {
			return false
		}
		return offers[i][0] < offers[j][0]
	})
	m := len(offers)
	dp := make([]int, m+1)
	ans := offers[0][2]
	for i := 0; i < m; i++ {
		if i == 0 {
			dp[i] = offers[i][2]
			continue
		}

		yes := offers[i][2]

		for j := i - 1; j >= 0; j-- {
			if offers[j][1] < offers[i][0] {
				yes = yes + dp[j]
				break
			}
		}
		dp[i] = max(dp[i-1], yes)
		ans = max(ans, dp[i])
	}

	return ans
}

type pair struct {
	start, offer int
}

func maximizeTheProfit(n int, offers [][]int) int {
	groups := make([][]pair, n)
	for _, ch := range offers {
		s, e, o := ch[0], ch[1], ch[2]
		groups[e] = append(groups[e], pair{s, o})
	}
	f := make([]int, n+1)
	for end, g := range groups {
		f[end+1] = f[end]
		for _, ch := range g {
			// 按正常的思维，dp方程应该是下面这个，但是这样的话会有-1的问题，解决办法就是整体前移一位
			// f[end] = max(f[end],f[ch.start-1]+ch.offer)
			f[end+1] = max(f[end+1], f[ch.start]+ch.offer)
		}
	}
	return f[n]
}
