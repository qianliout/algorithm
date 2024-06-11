package main

import (
	"container/heap"

	. "outback/algorithm/common/commonHeap"
)

func main() {

}

type KthLargest struct {
	Mh MinHeap
	K  int
}

func Constructor(k int, nums []int) KthLargest {
	mh := make(MinHeap, 0)
	for i := range nums {
		heap.Push(&mh, nums[i])
	}
	for mh.Len() > k {
		heap.Pop(&mh)
	}
	return KthLargest{Mh: mh, K: k}
}

func (this *KthLargest) Add(n int) int {
	heap.Push(&this.Mh, n)
	for this.Mh.Len() > this.K {
		heap.Pop(&this.Mh)
	}

	return this.Mh[0]
}
