package main

import (
	"container/heap"
	"sort"

	. "outback/algorithm/common/commonHeap"
)

func main() {

}

type KthLargest struct {
	MinHeap MaxHeap
	K       int
}

func Constructor(k int, nums []int) KthLargest {
	kk := KthLargest{
		MinHeap: make(MaxHeap, 0),
		K:       k,
	}

	sort.Ints(nums)
	for _, ch := range nums {
		heap.Push(&kk.MinHeap, ch)
		if len(kk.MinHeap) > k {
			heap.Pop(&kk.MinHeap)
		}
	}

	return kk
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.MinHeap, val)
	if len(this.MinHeap) > this.K {
		heap.Pop(&this.MinHeap)
	}
	return this.MinHeap[0]

}
