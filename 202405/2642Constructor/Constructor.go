package main

import (
	"math"
)

func main() {

}

type Graph struct {
	Edges [][]int
	G     [][]int
	N     int
}

func Constructor(n int, edges [][]int) Graph {
	gg := Graph{
		Edges: edges,
		G:     GenG(edges, n),
		N:     n,
	}
	return gg
}

func (this *Graph) AddEdge(edge []int) {
	this.N++
	this.Edges = append(this.Edges, edge)
	this.G = GenG(this.Edges, this.N)
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	return GetMin(this.G, this.N, node1, node2)
}

func GenG(edges [][]int, n int) [][]int {
	// 下标统一都做减一操作
	// 注意点1：这里最好不定义成 math.MaxInt32,因为下面有加法，可能会有溢出
	inf := math.MaxInt32
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, ch := range edges {
		x, y, z := ch[0], ch[1], ch[2]
		g[x][y] = z
	}
	return g
}

func GetMin(g [][]int, n, start, end int) int {
	// 下标统一都做减一操作
	// 注意点1：这里最好不定义成 math.MaxInt32,因为下面有加法，可能会有溢出
	inf := math.MaxInt32
	dis := make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[start] = 0
	done := make([]bool, n)
	for {
		x := -1
		for i, ok := range done {
			// 这里用 dis[i]<=dis[x] 还是用 dis[i]<dis[x]都能得到正确结果，还是还是推荐使用<=
			if !ok && (x < 0 || dis[i] <= dis[x]) {
				x = i
			}
		}

		if x < 0 { // 说明所有 n 个元素都更新的完了
			break
		}
		// 这里最好是>= 因为下面更新 dis[y] 时没有做判断，是直接加的，可能会比 inf 大
		// 即使不做判断，也不会有值超过 inf 因为是用的 min()操作
		// 说是不可达了
		if dis[x] >= inf {
			break
		}
		// 对于第一次的循环，x 一定是起始点,所以在这里更新 done 数组
		done[x] = true
		for y, d := range g[x] {
			// 这里可以做一步判断，判断是否 >= inf,也可以不判断，因为上面 >= inf 就都认为不可达
			// 也可以不判断
			if d >= inf {
				continue
			}
			dis[y] = min(dis[y], dis[x]+d)
		}
	}
	if dis[end] >= inf {
		return -1
	}
	return dis[end]
}
