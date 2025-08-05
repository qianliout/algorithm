package main

import (
	"math"
)

func main() {

}

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	nums := make([][]node, n)
	for _, ch := range edges {
		u, v, cnt := ch[0], ch[1], ch[2]
		// 一条边被细分为 cnt 个新节点时，实际需要经过 cnt + 1 条边
		// 例如：边 [u, v, 2] 会被细分为 u -> x1 -> x2 -> v，需要经过3条边
		nums[u] = append(nums[u], node{x: v, d: cnt + 1})
		nums[v] = append(nums[v], node{x: u, d: cnt + 1})
	}
	dis := dijkstra(nums, 0, n-1)
	ans := 0
	for _, d := range dis {
		if d <= maxMoves {
			ans++
		}
	}
	for _, ch := range edges {
		u, v, cnt := ch[0], ch[1], ch[2]
		a := max(0, maxMoves-dis[u])
		b := max(0, maxMoves-dis[v])
		ans += min(a+b, cnt)
	}
	return ans
}

func dijkstra(nums [][]node, start, end int) []int {
	inf := math.MaxInt / 10
	dis := make([]int, end+1)
	visit := make([]bool, end+1)
	for i := range dis {
		dis[i] = inf
	}
	dis[start] = 0
	// 不能有这一步
	// visit[0] = true
	for {
		x := -1
		minDis := inf
		for i := 0; i <= end; i++ {
			if !visit[i] && dis[i] < minDis {
				x = i
				minDis = dis[i]
			}
		}
		if x == -1 {
			break
		}
		visit[x] = true // 访问过了

		for _, no := range nums[x] {
			dis[no.x] = min(dis[no.x], dis[x]+no.d)
		}
	}
	return dis
}

type node struct {
	x, d int
}

/*
给你一个无向图（原始图），图中有 n 个节点，编号从 0 到 n - 1 。你决定将图中的每条边 细分 为一条节点链，每条边之间的新节点数各不相同。
图用由边组成的二维数组 edges 表示，其中 edges[i] = [ui, vi, cnti] 表示原始图中节点 ui 和 vi 之间存在一条边，cnti 是将边 细分 后的新节点总数。注意，cnti == 0 表示边不可细分。
要 细分 边 [ui, vi] ，需要将其替换为 (cnti + 1) 条新边，和 cnti 个新节点。新节点为 x1, x2, ..., xcnti ，新边为 [ui, x1], [x1, x2], [x2, x3], ..., [xcnti-1, xcnti], [xcnti, vi] 。
现在得到一个 新的细分图 ，请你计算从节点 0 出发，可以到达多少个节点？如果节点间距离是 maxMoves 或更少，则视为 可以到达 。
给你原始图和 maxMoves ，返回 新的细分图中从节点 0 出发 可到达的节点数
*/
