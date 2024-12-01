package main

import (
	"fmt"
)

func main() {
	fmt.Println(criticalConnections(4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}))
}

func criticalConnections(n int, connections [][]int) [][]int {
	g := make([][]int, n)
	for _, ch := range connections {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	ans := make([][]int, 0)
	ids := make([]int, n)
	for i := range ids {
		ids[i] = -1
	}
	var dfs func(i, fa, id int) int
	dfs = func(i, fa, id int) int {
		ids[i] = id

		for _, j := range g[i] {
			if j == fa {
				continue
			} else if ids[j] == -1 {
				ids[i] = min(ids[i], dfs(j, i, id+1))
			} else {
				ids[i] = min(ids[i], ids[j])
			}
		}
		if ids[i] == id && fa != -1 {
			ans = append(ans, []int{i, fa})
		}
		return ids[i]
	}
	dfs(0, -1, 0)
	return ans
}

// 中心思想，环中的链接都不是关键链接
// 环外的链接都是关键链接
