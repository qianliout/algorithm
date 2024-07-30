package main

import (
	"container/heap"

	. "outback/algorithm/common/commonHeap"
)

func main() {

}

func minStoneSum(piles []int, k int) int {
	hp := make(MaxHeap, 0)
	all := 0
	for _, ch := range piles {
		all += ch
		heap.Push(&hp, ch)
	}

	for k > 0 && hp.Len() > 0 {
		pop := heap.Pop(&hp).(int)
		all -= pop / 2
		if pop-pop/2 > 0 {
			heap.Push(&hp, pop-pop/2)
		}
		k--
	}
	return all
}
