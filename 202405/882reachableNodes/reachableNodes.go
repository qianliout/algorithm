package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(reachableNodes([][]int{{0, 1, 10}, {0, 2, 1}, {1, 2, 2}}, 6, 3))
}

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	g := make([][]neighbor, n)
	for _, ch := range edges {
		x, y, z := ch[0], ch[1], ch[2]+1
		g[x] = append(g[x], neighbor{y, z})
		g[y] = append(g[y], neighbor{x, z})
	}
	dist := dijkstra(g, 0, n-1)
	ans := 0
	for _, d := range dist {
		if d <= maxMoves {
			ans++
		}
	}
	for _, ch := range edges {
		x, y, z := ch[0], ch[1], ch[2]
		a := max(maxMoves-dist[x], 0)
		b := max(maxMoves-dist[y], 0)
		ans += min(a+b, z)
	}
	return ans
}

type neighbor struct {
	x, d int
}

func dijkstra(g [][]neighbor, start int, end int) []int {
	inf := math.MaxInt
	hp := make(PriorityQueue, 0)
	heap.Push(&hp, &IntItem{Value: 0, Priority: 0})
	dis := make([]int, end+1)
	for i := range dis {
		dis[i] = inf
	}
	dis[start] = 0
	for hp.Len() > 0 {
		po := heap.Pop(&hp).(*IntItem)
		for _, next := range g[po.Value] {
			newDis := dis[po.Value] + next.d
			if dis[next.x] > newDis {
				dis[next.x] = newDis
				heap.Push(&hp, &IntItem{Value: next.x, Priority: newDis})
			}
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
