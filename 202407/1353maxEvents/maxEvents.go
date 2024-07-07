package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"

	. "outback/algorithm/common/commonHeap"
)

func main() {
	fmt.Println(maxEvents([][]int{{1, 2}, {1, 2}, {3, 3}, {1, 5}, {1, 5}}))
}

// 按结束时间排序，会超时
func maxEvents1(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		if events[i][1] < events[j][1] {
			return true
		} else if events[i][1] > events[j][1] {
			return false
		}
		return events[i][0] < events[j][0]
	})

	used := make(map[int]bool)
	n := len(events)
	for i := 0; i < n; i++ {
		start, end := events[i][0], events[i][1]
		for k := start; k <= end; k++ {
			if !used[k] {
				used[k] = true
				break
			}
		}
	}
	return len(used)
}

func maxEvents(events [][]int) int {
	sort.SliceStable(events, func(i, j int) bool { return events[i][0] < events[j][0] })
	h := make(MinHeap, 0)
	end := int(math.Pow10(5))
	count := 0
	for t := 0; t <= end; t++ {
		for len(events) > 0 {
			if t >= events[0][0] {
				heap.Push(&h, events[0][1])
				events = events[1:]
				continue
			}
			break
		}
		for h.Len() > 0 && t > h[0] {
			heap.Pop(&h)
		}
		if h.Len() > 0 {
			pop := heap.Pop(&h).(int)
			if t <= pop {
				count++
			}
		}
	}
	return count
}
