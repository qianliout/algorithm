package main

import (
	"container/heap"
)

func main() {

}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	capitalHp := make(CapitalHeap, 0)
	for i, ch := range profits {
		heap.Push(&capitalHp, pair{profit: ch, capital: capital[i]})
	}
	profitHp := make(ProfitHeap, 0)
	for k > 0 {
		for capitalHp.Len() > 0 {
			pop := heap.Pop(&capitalHp).(pair)
			if pop.capital > w {
				heap.Push(&capitalHp, pop)
				break
			}
			heap.Push(&profitHp, pop)
		}
		if profitHp.Len() == 0 {
			break
		}
		pop := heap.Pop(&profitHp).(pair)
		w += pop.profit
		k--
	}
	return w
}

type pair struct {
	profit  int
	capital int
}

type CapitalHeap []pair

func (h CapitalHeap) Len() int           { return len(h) }
func (h CapitalHeap) Less(i, j int) bool { return h[i].capital < h[j].capital }
func (h CapitalHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *CapitalHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *CapitalHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *CapitalHeap) Peek() interface{} {
	if len(*h) > 0 {
		return (*h)[0]
	}
	return 0
}

type ProfitHeap []pair

func (h ProfitHeap) Len() int           { return len(h) }
func (h ProfitHeap) Less(i, j int) bool { return h[i].profit > h[j].profit }
func (h ProfitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ProfitHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *ProfitHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *ProfitHeap) Peek() interface{} {
	if len(*h) > 0 {
		return (*h)[0]
	}
	return 0
}
