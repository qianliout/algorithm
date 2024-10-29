package dijkstra

import (
	"container/heap"
	"math"
)

// https://leetcode.cn/problems/reachable-nodes-in-subdivided-graph/description/

type neighbor struct {
	x, d int
}

func dijkstra(g [][]neighbor, start int, end int) []int {
	inf := math.MaxInt
	hp := make(MinHeap, 0)
	heap.Push(&hp, Item{X: 0, D: 0})
	dis := make([]int, end+1)
	for i := range dis {
		dis[i] = inf
	}
	dis[start] = 0
	for hp.Len() > 0 {
		po := heap.Pop(&hp).(Item)
		for _, next := range g[po.X] {
			newDis := dis[po.D] + next.d
			if dis[next.x] > newDis {
				dis[next.x] = newDis
				heap.Push(&hp, Item{X: next.x, D: newDis})
			}
		}
	}
	return dis
}

type Item struct {
	X int
	D int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].D < h[j].D }
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
