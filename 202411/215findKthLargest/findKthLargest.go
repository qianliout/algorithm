package main

import (
	"container/heap"

	. "outback/algorithm/common/commonHeap"
)

func main() {

}

func findKthLargest(nums []int, k int) int {
	hm := make(MinHeap, 0)
	for _, ch := range nums {
		heap.Push(&hm, ch)
		if hm.Len() > k {
			heap.Pop(&hm)
		}
	}
	return hm[0]
}
