package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(maxAltitude([]int{14, 2, 27, -5, 28, 13, 39}, 3))
}

func maxAltitude(heights []int, limit int) []int {
	heapMx := make(MaxHeap, 0)
	n := len(heights)
	ans := make([]int, 0)
	for ri := 0; ri < n; ri++ {
		heap.Push(&heapMx, Node{idx: ri, value: heights[ri]})
		if ri >= limit-1 {
			for heapMx.Len() > 0 {
				pike := heapMx[0]
				if pike.idx <= ri-limit {
					heap.Pop(&heapMx)
					continue
				}
				ans = append(ans, pike.value)
				break
			}
		}
	}

	return ans
}

type Node struct {
	idx   int
	value int
}
type MaxHeap []Node

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {

	if h[i].value != h[j].value {
		return h[i].value > h[j].value
	}
	return h[i].idx < h[j].idx
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Peek() interface{} {
	if len(*h) > 0 {
		return (*h)[0]
	}
	return 0
}
