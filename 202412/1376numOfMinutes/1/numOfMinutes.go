package main

import (
	"fmt"
)

func main() {
	fmt.Println(numOfMinutes(6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0}))
	fmt.Println(numOfMinutes(7, 6, []int{1, 2, 3, 4, 5, 6, -1}, []int{0, 6, 5, 4, 3, 2, 1}))
}

// bfs 是不好做的
// 其实就是树的最大深度的变体
func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		if manager[i] >= 0 {
			g[manager[i]] = append(g[manager[i]], i)
		}
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 || i >= n {
			return 0
		}
		mx := 0
		for _, nx := range g[i] {
			mx = max(mx, dfs(nx)+informTime[i])
		}
		return mx
	}
	return dfs(headID)
}
