package main

import (
	"math"
)

func main() {

}

/*
有 n 个城市，按从 0 到 n-1 编号。给你一个边数组 edges，其中 edges[i] = [fromi, toi, weighti] 代表 fromi 和 toi 两个城市之间的双向加权边，距离阈值是一个整数 distanceThreshold。
返回能通过某些路径到达其他城市数目最少、且路径距离 最大 为 distanceThreshold 的城市。如果有多个这样的城市，则返回编号最大的城市。
注意，连接城市 i 和 j 的路径的距离等于沿该路径的所有边的权重之和。
*/
func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	inf := math.MaxInt / 22
	g := make([][]int, n)
	mem := make([][][]int, n)
	for i := range mem {
		mem[i] = make([][]int, n)
		for j := range mem[i] {
			mem[i][j] = make([]int, n)
			for k := range mem[i][j] {
				mem[i][j][k] = -1
			}
		}
	}

	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}

	for _, ch := range edges {
		fr, to, w := ch[0], ch[1], ch[2]
		g[fr][to] = w
		g[to][fr] = w
	}

	minCnt := n + 10
	ans := -1

	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if j != i && distanceThreshold >= dfs(g, n-1, i, j, mem) {
				cnt++
			}
		}
		if minCnt >= cnt {
			minCnt = cnt
			ans = i
		}
	}
	return ans
}

// res := min(dfs(k-1, i, j), dfs(k-1, i, k)+dfs(k-1, k, j))
/*
	从i到j 是否经过 k，这里k,可以认为是 i--j 中的最大值
*/
func dfs(g [][]int, k, i, j int, mem [][][]int) int {
	if mem[k][i][j] >= 0 {
		return mem[k][i][j]
	}

	if k < 0 {
		mem[k][i][j] = g[i][j]
		return g[i][j]
	}
	d := min(dfs(g, k-1, i, j, mem), dfs(g, k-1, i, k, mem)+dfs(g, k-1, k, j, mem))
	mem[k][i][j] = d
	return d
}
