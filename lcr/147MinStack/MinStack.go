package main

import (
	"container/heap"
	. "outback/algorithm/common/commonHeap"
)

func main() {
}

type MinStack struct {
	St     []int
	Exit   map[int]int
	HeadMi MinHeap
}

func Constructor() MinStack {
	s := MinStack{
		St:     make([]int, 0),
		Exit:   make(map[int]int),
		HeadMi: make(MinHeap, 0),
	}
	return s
}

func (ms *MinStack) Push(x int) {
	ms.St = append(ms.St, x)
	ms.Exit[x]++
	heap.Push(&ms.HeadMi, x)
}

func (ms *MinStack) Pop() {
	if len(ms.St) == 0 {
		return
	}
	last := ms.St[len(ms.St)-1]
	ms.St = ms.St[:len(ms.St)-1]
	ms.Exit[last]--
}

func (ms *MinStack) Top() int {
	if len(ms.St) == 0 {
		return -1
	}
	return ms.St[len(ms.St)-1]
}

func (ms *MinStack) GetMin() int {
	for ms.HeadMi.Len() > 0 {
		pik := ms.HeadMi[0]
		if ms.Exit[pik] <= 0 {
			heap.Pop(&ms.HeadMi)
			continue
		}
		return pik
	}
	return -1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
