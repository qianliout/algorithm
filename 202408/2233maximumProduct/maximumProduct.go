package main

import (
	"container/heap"
	"math"
)

func main() {

}

func maximumProduct(nums []int, k int) int {
	mh := make(MinHeap, 0)
	mod := int(math.Pow10(9)) + 7
	for _, ch := range nums {
		heap.Push(&mh, ch)
	}
	for k > 0 {
		pop := heap.Pop(&mh).(int)
		heap.Push(&mh, pop+1)
		k--
	}
	all := 1
	for len(mh) > 0 {
		pop := heap.Pop(&mh).(int)
		all = (all * pop) % mod
	}
	return all
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] <= h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
