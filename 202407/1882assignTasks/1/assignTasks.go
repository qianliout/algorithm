package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// fmt.Println(assignTasks([]int{3, 3, 2}, []int{1, 2, 3, 2, 1, 2}))
	fmt.Println(assignTasks([]int{10, 63, 95, 16, 85, 57, 83, 95, 6, 29, 71}, []int{70, 31, 83, 15, 32, 67, 98, 65, 56, 48, 38, 90, 5})) // 8,0,3,9,5,1,10,6,4,2,7,9,0
}

// 只用一个堆，搞不定
func assignTasks(servers []int, tasks []int) []int {
	n := len(tasks)
	pq := make(PriorityQueue, 0)
	ans := make([]int, n)
	for i, ch := range servers {
		it := &Item{
			Idx:       i,
			Value:     ch,
			EnqueueTi: 0,
		}
		heap.Push(&pq, it)
	}

	// at, i := 0, 0
	// 如果没有空闲服务器，则必须等待，直到出现一台空闲服务器，并 尽可能早 地处理剩余任务。 如果有多项任务等待分配，则按照 下标递增 的顺序完成分配。
	// for i < n {
	// 	if pq.Len() == 0 {
	//
	// 	}
	//
	// }

	for i, ch := range tasks {
		find := make([]*Item, 0)

		for pq.Len() > 0 && pq[0].EnqueueTi > i {
			pop := heap.Pop(&pq).(*Item)
			find = append(find, pop)
		}
		// 如果没有空闲服务器，则必须等待，直到出现一台空闲服务器，并 尽可能早 地处理剩余任务。 如果有多项任务等待分配，则按照 下标递增 的顺序完成分配。
		if pq.Len() == 0 {
			find = change(find, i)
			for _, it := range find {
				heap.Push(&pq, it)
			}
			find = make([]*Item, 0)
		}
		pop := heap.Pop(&pq).(*Item)

		ans[i] = pop.Idx
		pop.EnqueueTi = i + ch

		find = append(find, pop)
		for _, it := range find {
			heap.Push(&pq, it)
		}
	}
	return ans
}

func change(ss []*Item, i int) []*Item {
	mi := ss[0].EnqueueTi
	for _, ch := range ss {
		mi = min(mi, ch.EnqueueTi)
	}
	for _, ch := range ss {
		ch.EnqueueTi -= mi - i
	}
	return ss
}

type Item struct {
	Idx       int
	Value     int // 权重
	EnqueueTi int // 可以开始工作的时间
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {

	if pq[i].Value < pq[j].Value {
		return true
	} else if pq[i].Value > pq[j].Value {
		return false
	} else {
		return pq[i].Idx < pq[j].Idx
	}
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
