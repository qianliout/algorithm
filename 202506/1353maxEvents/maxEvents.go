package main

import (
	"container/heap"
	"sort"

	. "outback/algorithm/common/commonHeap"
)

func main() {

}

func maxEvents(events [][]int) int {
	sort.SliceStable(events, func(i, j int) bool { return events[i][0] < events[j][0] })
	h := make(MinHeap, 0)
	end := 100000
	ans := 0
	for e := 0; e <= end; e++ {
		for len(events) > 0 {
			if e >= events[0][0] {
				heap.Push(&h, events[0][1])
				events = events[1:]
				continue
			}
			break
		}
		for h.Len() > 0 && e > h[0] {
			heap.Pop(&h)
		}
		if h.Len() > 0 {
			pop := heap.Pop(&h).(int)
			if e <= pop {
				ans++
			}
		}
	}
	return ans
}
