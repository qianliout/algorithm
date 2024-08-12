package main

import (
	"container/heap"
)

func main() {

}

func totalCost(costs []int, k int, c int) int64 {
	mh := make(MinHeap, 0)
	n := len(costs)
	left, right := c, n-c-1
	for i := 0; i < c; i++ {
		heap.Push(&mh, pair{id: i, cost: costs[i]})
	}
	for i := n - 1; i >= max(0, left, right+1); i-- {
		heap.Push(&mh, pair{id: i, cost: costs[i]})
	}
	ans := int64(0)
	for i := 0; i < k; i++ {
		pop := heap.Pop(&mh).(pair)
		if pop.id > right {
			if right >= left {
				heap.Push(&mh, pair{id: right, cost: costs[right]})
			}
			right--
		} else {
			if right >= left {
				heap.Push(&mh, pair{id: left, cost: costs[left]})
			}
			left++
		}
		ans += int64(pop.cost)
	}

	return ans
}

type pair struct {
	id   int
	cost int
}

type MinHeap []pair

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].cost != h[j].cost {
		return h[i].cost < h[j].cost
	}
	return h[i].id < h[j].id
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
