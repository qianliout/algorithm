package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(halveArray([]int{5, 19, 8, 1}))

}

func halveArray(nums []int) int {
	mp := make(MaxHeap, 0)
	sum := 0
	for _, ch := range nums {
		sum += ch
		heap.Push(&mp, float64(ch))
	}
	var sub float64
	ans := 0
	for sub < float64(sum)/float64(2) {
		pop := heap.Pop(&mp).(float64)
		ans++
		sub += pop / 2
		heap.Push(&mp, pop/2)
	}
	return ans
}

type MaxHeap []float64

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
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
