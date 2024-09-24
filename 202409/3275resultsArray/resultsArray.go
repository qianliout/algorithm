package main

import (
	"container/heap"

	. "outback/algorithm/common/commonHeap"
)

func main() {

}

func resultsArray(queries [][]int, k int) []int {
	n := len(queries)
	h := make(MaxHeap, 0)
	ans := make([]int, n)
	for i, q := range queries {
		x := abs(q[0]) + abs(q[1])
		heap.Push(&h, x)
		if h.Len() > k {
			heap.Pop(&h)
		}
		if h.Len() == k {
			ans[i] = h.Peek().(int)
		} else {
			ans[i] = -1
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
