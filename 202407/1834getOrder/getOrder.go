package main

import (
	"container/heap"
	"sort"
)

func main() {

}

func getOrder(tasks [][]int) []int {
	items := make([]*Item, 0)
	for i, ch := range tasks {
		it := &Item{
			Idx:          i,
			EnqueueTi:    ch[0],
			ProcessingTi: ch[1],
		}
		items = append(items, it)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].EnqueueTi <= items[j].EnqueueTi })
	pq := make(PriorityQueue, 0)
	at := 1
	n := len(items)
	ans := make([]int, 0)
	// for i, j := 0, 0; i < n; i++ {
	// 	if pq.Len() == 0 {
	// 		at = max(at, items[j].EnqueueTi)
	// 	}
	// 	for j < n && items[j].EnqueueTi <= at {
	// 		heap.Push(&pq, items[j])
	// 		j++
	// 	}
	// 	pop := heap.Pop(&pq).(*Item)
	// 	ans = append(ans, pop.Idx)
	// 	at += pop.ProcessingTi
	// }
	// 这样写判断会更容易理解一点
	j := 0
	for j < n || pq.Len() > 0 {
		if pq.Len() == 0 && j < n {
			at = max(at, items[j].EnqueueTi)
		}
		for j < n && items[j].EnqueueTi <= at {
			heap.Push(&pq, items[j])
			j++
		}
		pop := heap.Pop(&pq).(*Item)
		ans = append(ans, pop.Idx)
		at += pop.ProcessingTi
	}

	return ans
}

type Item struct {
	Idx          int
	EnqueueTi    int
	ProcessingTi int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].ProcessingTi < pq[j].ProcessingTi {
		return true
	} else if pq[i].ProcessingTi > pq[j].ProcessingTi {
		return false
	}
	return pq[i].Idx < pq[j].Idx
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
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
