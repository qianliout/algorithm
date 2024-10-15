package main

import (
	"container/heap"
)

func main() {

}

func minimumTime(n int, edges [][]int, disappear []int) []int {
	g := make([][]pair, n)
	for _, ch := range edges {
		x, y, l := ch[0], ch[1], ch[2]
		g[x] = append(g[x], pair{to: y, le: l})
		g[y] = append(g[y], pair{to: x, le: l})
	}
	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	dis[0] = 0
	hq := make(MinHeap, 0)
	heap.Push(&hq, distance{node: 0, dis: 0})
	// 不用堆优化了，会超时
	for len(hq) > 0 {
		qu := heap.Pop(&hq).(distance)
		no, di := qu.node, qu.dis
		if dis[no] >= 0 && dis[no] < di {
			continue
		}
		for _, ch := range g[no] {
			y, wt := ch.to, ch.le
			nd := qu.dis + wt
			if nd < disappear[y] && (dis[y] < 0 || dis[y] > nd) {
				dis[y] = nd
				heap.Push(&hq, distance{node: y, dis: nd})
			}
		}
	}
	return dis
}

type pair struct {
	to, le int
}

type distance struct {
	node, dis int
}

type MinHeap []distance

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].dis <= h[j].dis }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(distance))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
