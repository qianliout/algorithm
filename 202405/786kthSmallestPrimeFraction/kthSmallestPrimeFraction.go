package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3))
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	hp := make(PriorityQueue, 0)
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			it := &IntItem{A: arr[i], B: arr[j], AD: i, BD: j}
			if hp.Len() < k {
				heap.Push(&hp, it)
				continue
			}

			if Less(it, hp[0]) {
				heap.Push(&hp, it)
				heap.Pop(&hp)
			}
		}
	}
	top := heap.Pop(&hp).(*IntItem)
	return []int{top.A, top.B}
}
func Less(i, j *IntItem) bool {
	return i.A*j.B < j.A*i.B
}

type IntItem struct {
	A, B   int
	AD, BD int
	// Value    int // The Key of the item; arbitrary.
	// Priority int // The Priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*IntItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].A*pq[j].B > pq[j].A*pq[i].B
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*IntItem)
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
