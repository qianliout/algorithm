package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(countGoodNodes([][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}}))

}

// 这样写是对的，但是会超过内存限制
func countGoodNodes1(edges [][]int) int {
	n := len(edges)
	g := make([][]int, n+1)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	mem := make([][]int, n+2)
	for i := range mem {
		mem[i] = make([]int, n+2)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs1 func(i, fa int) int
	dfs1 = func(x, fa int) int {
		if x > n || x < 0 {
			return 0
		}
		if mem[x][fa] != -1 {
			return mem[x][fa]
		}
		cnt := 1
		for _, y := range g[x] {
			if y != fa {
				cnt += dfs1(y, x)
			}
		}
		mem[x][fa] = cnt
		return cnt
	}
	ans := 0

	var dfs2 func(x, fa int)
	dfs2 = func(x, fa int) {
		if x > n || x < 0 {
			return
		}
		child := make([]int, 0)
		for _, y := range g[x] {
			if y != fa {
				dfs2(y, x)
				cnt := dfs1(y, x)
				child = append(child, cnt)
			}
		}
		if len(slices.Compact(child)) <= 1 {
			ans++
		}
	}
	dfs2(0, n+1)

	return ans
}

func countGoodNodes(edges [][]int) int {
	n := len(edges)
	g := make([][]int, n+1)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	// 当前节点的子为根的树大小
	// 从不同的父节点走过来，树的大小可能不一样，但是dfs2中已明确了父节点的路径，所以只记录一个节点能用
	mem := make([]int, n+2)
	for i := range mem {
		mem[i] = -1
	}
	var dfs1 func(i, fa int) int
	dfs1 = func(x, fa int) int {
		if x > n || x < 0 {
			return 0
		}
		if mem[x] != -1 {
			return mem[x]
		}
		cnt := 1
		for _, y := range g[x] {
			if y != fa {
				cnt += dfs1(y, x)
			}
		}
		mem[x] = cnt
		return cnt
	}
	ans := 0

	var dfs2 func(x, fa int)
	dfs2 = func(x, fa int) {
		if x > n || x < 0 {
			return
		}
		child := make([]int, 0)
		for _, y := range g[x] {
			if y != fa {
				dfs2(y, x)
				cnt := dfs1(y, x)
				child = append(child, cnt)
			}
		}
		if len(slices.Compact(child)) <= 1 {
			ans++
		}
	}
	dfs2(0, n+1)

	return ans
}
