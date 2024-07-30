package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(kthLargestNumber([]string{"100", "3", "6", "7", "10", "2"}, 1))
	fmt.Println(kthLargestNumber([]string{"8", "4676", "30", "7", "38"}, 4))
}

func kthLargestNumber(nums []string, k int) string {
	// 返回 nums 中表示第 k 大整数的字符串。
	pq := make(MinHeap, 0)
	for _, ch := range nums {
		if pq.Len() < k {
			heap.Push(&pq, ch)
			continue
		}
		heap.Push(&pq, ch)
		heap.Pop(&pq)
	}
	return pq[0]
}

type MinHeap []string

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	a, b := h[i], h[j]
	if len(a) < len(b) {
		return true
	} else if len(a) > len(b) {
		return false
	}
	return a <= b
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(string))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
