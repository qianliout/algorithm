package main

import (
	"container/heap"

	. "outback/algorithm/common/commonHeap"
)

func main() {

}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	minH := make(MinHeap, 0)
	n := len(heights)
	b := 0
	for i := 0; i < n-1; i++ {
		h := heights[i+1] - heights[i]
		if h <= 0 {
			continue
		}
		heap.Push(&minH, h)
		if minH.Len() > ladders {
			p := heap.Pop(&minH).(int)
			b += p
		}
		if b > bricks {
			return i
		}
	}
	return n - 1
}
