package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(secondMinimum(5, [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}, 3, 5))
	fmt.Println(secondMinimum(2, [][]int{{1, 2}}, 3, 2))
}

func secondMinimum(n int, edges [][]int, ti int, change int) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0]-1, e[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	inf := 1 << 32
	dis1 := make([]int, n)
	dis2 := make([]int, n)
	for i := range dis1 {
		dis1[i] = inf
		dis2[i] = inf
	}
	dis1[0] = 0

	hm := make(MinHeap, 0)
	heap.Push(&hm, IntItem{Step: 0, Idx: 0})
	// 要注意的是，只会有一条最短路，所以不用 visit 数组
	for hm.Len() > 0 {
		pop := heap.Pop(&hm).(IntItem)
		x := pop.Idx
		d := pop.Step

		if d > dis2[x] {
			continue
		}
		for _, y := range g[x] {
			nd := d + ti
			// 动态边权
			if (d/change)&1 == 1 {
				// 说明是红灯
				nd = d + ti + change - d%change
			}
			if nd < dis2[y] && nd > dis1[y] {
				dis2[y] = nd
				heap.Push(&hm, IntItem{Step: nd, Idx: y})
			}
			if nd < dis1[y] {
				dis2[y] = dis1[y]
				dis1[y] = nd
				heap.Push(&hm, IntItem{Step: nd, Idx: y})
			}
		}
	}
	return dis2[n-1]
}

type MinHeap []IntItem

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Step <= h[j].Step }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(IntItem))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type IntItem struct {
	Step int
	Idx  int
}
