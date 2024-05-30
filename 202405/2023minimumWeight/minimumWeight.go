package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	// fmt.Println(minimumWeight(3, [][]int{{0, 1, 2}, {2, 1, 1}}, 0, 1, 2))
	fmt.Println(minimumWeight(6, [][]int{{0, 2, 2}, {0, 5, 6}, {1, 0, 3}, {1, 4, 5}, {2, 1, 1}, {2, 3, 3}, {2, 3, 4}, {3, 4, 2}, {4, 5, 1}}, 0, 1, 5))
}

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	// 带权有向 图的节点数
	g := make([][]neighbor, n)
	rg := make([][]neighbor, n)
	for _, ch := range edges {
		x, y, z := ch[0], ch[1], ch[2]
		g[x] = append(g[x], neighbor{x: y, d: z})
		rg[y] = append(rg[y], neighbor{x: x, d: z})
	}
	d1 := dijkstra(g, src1, n)
	d2 := dijkstra(g, src2, n)
	d3 := dijkstra(rg, dest, n)
	ans := math.MaxInt / 3
	for i := 0; i < n; i++ {
		if d1[i] == -1 || d2[i] == -1 || d3[i] == -1 {
			continue
		}

		ans = min(ans, d1[i]+d2[i]+d3[i])
	}
	if ans >= math.MaxInt/3 {
		return -1
	}
	return int64(ans)
}

type neighbor struct {
	x, d int
}

func dijkstra(g [][]neighbor, start int, n int) []int {
	inf := math.MaxInt / 3
	hp := make(PriorityQueue, 0)
	heap.Push(&hp, &IntItem{Value: start, Priority: 0})
	dis := make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[start] = 0
	for hp.Len() > 0 {
		po := heap.Pop(&hp).(*IntItem)
		// 这一步可以优化速度
		// 说明这个 PO是之前加入的，但是 后上一步的操作中，dis[po.Value]已被其他邻居更新了，所以这个距离就是一个无效的距离了
		if po.Priority > dis[po.Value] {
			continue
		}
		for _, next := range g[po.Value] {
			newDis := dis[po.Value] + next.d
			if dis[next.x] > newDis {
				dis[next.x] = newDis
				heap.Push(&hp, &IntItem{Value: next.x, Priority: newDis})
			}
		}
	}
	for i := range dis {
		if dis[i] >= inf {
			dis[i] = -1
		}
	}

	return dis
}

type IntItem struct {
	Value    int // The Key of the item; arbitrary.
	Priority int // The Priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*IntItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, Priority so we use greater than here.
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*IntItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
