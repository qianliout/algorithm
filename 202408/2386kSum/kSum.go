package main

import (
	"container/heap"
	"sort"
)

func main() {

}

func kSum(nums []int, k int) int64 {
	mh := make(MinHeap, 0)
	all := 0
	for i, ch := range nums {
		if ch > 0 {
			all += ch
		} else {
			nums[i] = -ch
		}
	}
	sort.Ints(nums)
	n := len(nums)
	heap.Push(&mh, pair{0, 0})
	for j := 0; j < k; j++ {
		pop := heap.Pop(&mh).(pair)
		i := pop.idx
		if i < n {
			heap.Push(&mh, pair{pop.sum + nums[i], i + 1})
			if i > 0 {
				heap.Push(&mh, pair{pop.sum + nums[i] - nums[i-1], i + 1})
			}
		}
	}
	return int64(all - mh[0].sum)
}

type pair struct {
	sum int
	idx int
}

type MinHeap []pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].sum <= h[j].sum }
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
