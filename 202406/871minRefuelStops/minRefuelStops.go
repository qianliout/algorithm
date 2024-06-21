package main

import (
	"container/heap"

	. "outback/algorithm/common/commonHeap"
)

func main() {

}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	ans, cur := 0, startFuel
	hp := make(MaxHeap, 0)
	i := 0 // 当前加油站的index
	for cur < target {
		if i < len(stations) && cur >= stations[i][0] {
			heap.Push(&hp, stations[i][1])
			i++
		} else {
			if len(hp) == 0 {
				return -1
			}
			pop := heap.Pop(&hp).(int)
			ans++
			cur += pop
		}
	}

	return ans
}
