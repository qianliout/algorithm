package main

import (
	"container/heap"
	"math"

	. "outback/algorithm/common/commonHeap"
)

func main() {

}

func pickGifts(gifts []int, k int) int64 {
	hp := make(MaxHeap, 0)
	sum := 0
	for _, ch := range gifts {
		heap.Push(&hp, ch)
		sum += ch
	}
	ans := 0
	for k > 0 {
		pop := heap.Pop(&hp).(int)
		n := int(math.Sqrt(float64(pop)))
		ans += pop - n
		heap.Push(&hp, n)
		k--
	}

	return int64(max(0, sum-ans))
}
