package main

import (
	"fmt"
)

func main() {
	fmt.Println(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 2, 4))
	fmt.Println(frogPosition(8, [][]int{{2, 1}, {3, 2}, {4, 1}, {5, 1}, {6, 4}, {7, 1}, {8, 7}}, 7, 7))
}

// 自底向上
func frogPosition1(n int, edges [][]int, t int, target int) float64 {
	g := make([][]int, n+1)
	g[1] = append(g[1], 0)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(x, fa int, leftT int) int
	dfs = func(x, fa int, leftT int) int {
		// if leftT < 0 {
		// 	return 0
		// }
		if leftT == 0 {
			if x == target { // 恰好到达
				return 1
			}
			return 0
		}
		if x == target {
			// target 是叶子，停在原地
			if len(g[x]) == 1 {
				return 1
			}
			return 0
		}
		for _, d := range g[x] {
			if d == fa {
				continue
			}
			find := dfs(d, x, leftT-1)
			if find > 0 {
				return find * (len(g[x]) - 1)
			}
		}
		return 0

	}
	find := dfs(1, 0, t)
	if find > 0 {
		return float64(1) / float64(find)
	}
	return 0
}

// 自订向下
func frogPosition(n int, edges [][]int, t int, target int) float64 {
	g := make([][]int, n+1)
	g[1] = append(g[1], 0)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var ans float64
	var dfs func(x, fa int, leftT int, prod int) bool
	dfs = func(x, fa int, leftT int, prod int) bool {

		// t 秒后必须在 target（恰好到达，或者 target 是叶子停在原地）
		if x == target && (leftT == 0 || len(g[x]) == 1) {
			ans = float64(1) / float64(prod)
			return true
		}
		// 上一个判断了恰好到达的情况，这里提前结束
		if x == target || leftT <= 0 {
			return false
		}

		for _, d := range g[x] {
			if d == fa {
				continue
			}
			if dfs(d, x, leftT-1, prod*(len(g[x])-1)) {
				return true
			}
		}
		return false
	}

	find := dfs(1, 0, t, 1)
	if find {
		return ans
	}
	return 0
}
