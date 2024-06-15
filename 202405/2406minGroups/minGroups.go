package main

import (
	"container/heap"
	"fmt"
	"sort"

	. "outback/algorithm/common/commonHeap"
)

func main() {
	fmt.Println(minGroups([][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}}))
}

// 这个题目要理解才得行，这样做是对的
func minGroups(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	hm := make(MinHeap, 0)
	for _, ch := range intervals {
		if hm.Len() == 0 || ch[0] <= hm[0] {
			heap.Push(&hm, ch[1])
		} else {
			hm[0] = max(hm[0], ch[1])
			heap.Fix(&hm, 0)
		}
	}
	return hm.Len()
}
