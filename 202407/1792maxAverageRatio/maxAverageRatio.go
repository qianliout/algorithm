package main

import (
	"container/heap"
)

func main() {

}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	pq := make(PriorityQueue, 0)

	for _, ch := range classes {
		pass, total := ch[0], ch[1]
		heap.Push(&pq, &IntItem{
			Pass:     pass,
			Total:    total,
			Priority: float64(pass+1)/float64(total+1) - float64(pass)/float64(total),
		})
	}
	for extraStudents > 0 {
		pop := heap.Pop(&pq).(*IntItem)

		it := &IntItem{Pass: pop.Pass + 1, Total: pop.Total + 1}
		it.Priority = float64(it.Pass+1)/float64(it.Total+1) - float64(it.Pass)/float64(it.Total)
		heap.Push(&pq, it)
		extraStudents--
	}
	// 求和了
	var ans float64
	for _, ch := range pq {
		ans += float64(ch.Pass) / float64(ch.Total)
	}

	return ans / float64(len(pq))
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*IntItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, Priority so we use greater than here.
	return pq[i].Priority > pq[j].Priority
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

type IntItem struct {
	Pass     int
	Total    int
	Priority float64
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}
