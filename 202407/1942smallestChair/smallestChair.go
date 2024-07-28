package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(smallestChair([][]int{{1, 4}, {2, 3}, {4, 6}}, 1))
}

func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	chair := make(MinHeap, 0)
	for i := 0; i < n; i++ {
		heap.Push(&chair, i)
	}
	changes := make(PriorityQueue, 0)

	for i, ch := range times {
		pa1 := &pair{ID: i, at: ch[0], flag: 1}
		pa2 := &pair{ID: i, at: ch[1], flag: 0}
		heap.Push(&changes, pa1)
		heap.Push(&changes, pa2)
	}

	occupiedChairs := make(map[int]int) // key是客人编号，value 是椅子编号
	ans := -1
	for changes.Len() > 0 && ans == -1 {
		p := heap.Pop(&changes).(*pair)
		if p.flag > 0 {
			// 说明需要占一个椅子了
			ch := heap.Pop(&chair).(int)
			occupiedChairs[p.ID] = ch
			if p.ID == targetFriend {
				ans = ch
				return ans
			}
		} else {
			// 说明这个椅子已用完了，需要归还
			ch := occupiedChairs[p.ID]
			heap.Push(&chair, ch)
		}
	}

	return 0
}

type pair struct {
	ID    int
	at    int
	flag  int // 1 表示占用，0表示没有占用
	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*pair

func (pq PriorityQueue) Len() int { return len(pq) }

// Less 这个判断是这个题目的关键
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].at != pq[j].at {
		return pq[i].at < pq[j].at
	} else {
		return pq[i].flag < pq[j].flag
	}
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pair)
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

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] <= h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
