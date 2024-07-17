package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(eatenApples([]int{1, 2, 3, 5, 2}, []int{3, 2, 1, 4, 2}))
	fmt.Println(eatenApples([]int{3, 0, 0, 0, 0, 2}, []int{3, 0, 0, 0, 0, 2}))
}

// days[i] 天后（也就是说，第 i + days[i] 天时）
func eatenApples(apples []int, days []int) int {
	hp := make(MinHeap, 0)
	n := len(apples)
	ans := 0
	ti := 0
	for ti < n || hp.Len() > 0 {
		if ti < n && apples[ti] > 0 {
			heap.Push(&hp, pair{Cnt: apples[ti], Ti: ti + days[ti] - 1})
		}

		for hp.Len() > 0 {
			pop := heap.Pop(&hp).(pair)
			if pop.Ti >= ti {
				ans++
				pop.Cnt--
				if pop.Cnt > 0 {
					heap.Push(&hp, pop)
				}
				break
			}
		}
		ti++
	}
	return ans
}

type pair struct {
	Cnt int // 数量
	Ti  int // 最后食用日期
}

type MinHeap []pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Ti <= h[j].Ti }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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
