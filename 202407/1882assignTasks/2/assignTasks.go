package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(assignTasks([]int{3, 3, 2}, []int{1, 2, 3, 2, 1, 2}))                                                                    // 2,2,0,2,1,2
	fmt.Println(assignTasks([]int{10, 63, 95, 16, 85, 57, 83, 95, 6, 29, 71}, []int{70, 31, 83, 15, 32, 67, 98, 65, 56, 48, 38, 90, 5})) // 8,0,3,9,5,1,10,6,4,2,7,9,0
}

// 只用一个堆，搞不定
func assignTasks(servers []int, tasks []int) []int {
	n := len(tasks)
	busy := make(PriorityQueueBusy, 0)
	idle := make(PriorityQueueIdle, 0)
	ans := make([]int, n)
	for i, ch := range servers {
		it := &Item{
			Idx:       i,
			Value:     ch,
			EnqueueTi: 0,
		}
		heap.Push(&idle, it)
	}
	at := 0 // 表示当前时间
	for i, ch := range tasks {
		// 如果当前没有空闲的服务器，那么就等到最先完成任务的服务器完成任务
		at = max(at, i)
		if idle.Len() == 0 {
			at = max(at, busy[0].EnqueueTi)
		}
		// 把已跑完任务的服务器加入到 idle 队列中
		for busy.Len() > 0 && busy[0].EnqueueTi <= at {
			pop := heap.Pop(&busy).(*Item)
			heap.Push(&idle, pop)
		}
		// 取一个空闲的服务器
		pop := heap.Pop(&idle).(*Item)
		ans[i] = pop.Idx
		pop.EnqueueTi = at + ch
		heap.Push(&busy, pop)
	}

	return ans
}

type Item struct {
	Idx       int
	Value     int // 权重
	EnqueueTi int // 可以开始工作的时间
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type PriorityQueueBusy []*Item

func (pq PriorityQueueBusy) Len() int { return len(pq) }

func (pq PriorityQueueBusy) Less(i, j int) bool {
	if pq[i].EnqueueTi < pq[j].EnqueueTi {
		return true
	} else if pq[i].EnqueueTi > pq[j].EnqueueTi {
		return false
	} else {
		return pq[i].Idx < pq[j].Idx
	}
}

func (pq PriorityQueueBusy) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueueBusy) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueueBusy) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type PriorityQueueIdle []*Item

func (pq PriorityQueueIdle) Len() int { return len(pq) }

func (pq PriorityQueueIdle) Less(i, j int) bool {
	if pq[i].Value < pq[j].Value {
		return true
	} else if pq[i].Value > pq[j].Value {
		return false
	} else {
		return pq[i].Idx < pq[j].Idx
	}
}

func (pq PriorityQueueIdle) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueueIdle) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueueIdle) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
