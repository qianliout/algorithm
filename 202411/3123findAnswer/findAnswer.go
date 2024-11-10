package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(findAnswer(6, [][]int{{0, 1, 4}, {0, 2, 1}, {1, 3, 2}, {1, 4, 3}, {1, 5, 1}, {2, 3, 1}, {3, 5, 3}, {4, 5, 2}}))
}

func findAnswer(n int, edges [][]int) []bool {
	g := make([][]pair, n)
	for i, ch := range edges {
		x, y, w := ch[0], ch[1], ch[2]
		g[x] = append(g[x], pair{i, y, w})
		g[y] = append(g[y], pair{i, x, w})
	}
	inf := 1 << 32
	dis := dijkstra(g)
	m := len(edges)
	ans := make([]bool, m)
	var dfs func(x int)
	if dis[n-1] == inf {
		return ans
	}
	vis := make([]bool, n)

	dfs = func(y int) {
		vis[y] = true
		for _, ch := range g[y] {
			x, w, i := ch.to, ch.w, ch.idx
			if dis[x]+w != dis[y] {
				continue
			}
			ans[i] = true
			if !vis[x] {
				dfs(x)
			}
		}
	}
	dfs(n - 1)
	return ans
}

type pair struct {
	idx, to, w int
}

func dijkstra(g [][]pair) []int {
	n := len(g)
	inf := 1 << 32
	dis := make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[0] = 0
	mh := make(MinHeap, 0)
	heap.Push(&mh, Item{X: 0, Dis: 0})
	for mh.Len() > 0 {
		pop := heap.Pop(&mh).(Item)
		x, d := pop.X, pop.Dis
		// 下面这个判断可有可没有
		// if d > dis[x] {
		// 	continue
		// }
		for _, ch := range g[x] {
			y, w, _ := ch.to, ch.w, ch.idx
			nd := d + w
			if nd < dis[y] {
				dis[y] = nd
				heap.Push(&mh, Item{X: y, Dis: nd})
			}
		}
	}
	return dis
}

type Item struct {
	X   int
	Dis int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Dis < h[j].Dis }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
