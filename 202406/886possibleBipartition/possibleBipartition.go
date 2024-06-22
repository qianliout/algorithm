package main

import (
	"fmt"
)

func main() {
	fmt.Println(possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}))
}

func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n)
	// for i := range g {
	// 	g[i] = make([]int, n)
	// }
	for _, ch := range dislikes {
		x, y := ch[0]-1, ch[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(i, c int) bool
	color := make([]int, n)
	// 两种颜色分别记为1，2,所以知道一种颜色c之后，另一种颜色可以记为 3-c
	dfs = func(i, c int) bool {
		if i >= n {
			return true
		}
		if color[i] != 0 && color[i] != c {
			return false
		}
		color[i] = c
		for _, nex := range g[i] {
			if color[nex] == c {
				return false
			}
			if color[nex] == 0 && !dfs(nex, 3-c) {
				return false
			}
		}
		return true
	}

	for i, c := range color {
		if c == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true
}
