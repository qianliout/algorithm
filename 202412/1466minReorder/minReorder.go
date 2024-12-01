package main

import (
	"fmt"
)

func main() {
	fmt.Println(minReorder(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}))
}

// 好理解的方法
func minReorder1(n int, connections [][]int) int {
	in := make([][]int, n) // 表示入
	out := make([][]int, n)
	for _, ch := range connections {
		x, y := ch[0], ch[1]
		in[y] = append(in[y], x)
		out[x] = append(out[x], y)
	}
	var dfs func(i int) int
	visit := make([]bool, n)
	visit[0] = true
	dfs = func(i int) int {
		ans := 0
		for _, nx := range out[i] {
			if visit[nx] {
				continue
			}
			visit[nx] = true
			ans += 1 + dfs(nx)
		}
		for _, nx := range in[i] {
			if visit[nx] {
				continue
			}
			visit[nx] = true
			ans += dfs(nx)
		}
		return ans
	}
	return dfs(0)
}

func minReorder(n int, connections [][]int) int {
	g := make([][]pair, n)
	for _, ch := range connections {
		a, b := ch[0], ch[1]
		g[a] = append(g[a], pair{b, 1})
		g[b] = append(g[b], pair{a, 0})
	}
	var dfs func(i, pa int) int

	dfs = func(i, pa int) int {
		res := 0
		for _, nex := range g[i] {
			if nex.b == pa {
				continue
			}
			res += nex.cost + dfs(nex.b, i)
		}
		return res
	}
	return dfs(0, -1)
}

type pair struct {
	b    int
	cost int // 有一条路，也就是不入边时，cost 等0，没有入边时 cost 等于1
}
