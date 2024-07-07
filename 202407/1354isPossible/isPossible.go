package main

import (
	"container/heap"
	"fmt"

	. "outback/algorithm/common/commonHeap"
)

func main() {
	fmt.Println(isPossible([]int{9, 3, 5}))
	fmt.Println(isPossible([]int{8, 5}))
}

func isPossible(target []int) bool {
	hm := make(MaxHeap, 0)
	sum := 0
	for _, ch := range target {
		heap.Push(&hm, ch)
		sum += ch
	}
	for len(hm) > 0 {
		curMax := heap.Pop(&hm).(int)
		if curMax == 1 {
			break
		}
		otherSum := sum - curMax
		if otherSum == 0 || curMax-otherSum < 1 {
			return false
		}
		tm := curMax % otherSum
		if tm == 0 {
			tm = otherSum
		}
		heap.Push(&hm, tm)
		sum = sum - curMax + tm
	}
	return true
}
