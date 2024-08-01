package main

import (
	"container/heap"
	"sort"
)

func main() {

}

func maxTwoEvents(events [][]int) int {
	hp := make(MinHeap, 0)

	sort.Slice(events, func(i, j int) bool { return events[i][0] < events[j][0] })
	ans := 0
	mx := 0
	for _, ch := range events {
		s, e, v := ch[0], ch[1], ch[2]
		for len(hp) > 0 && hp[0].end < s {
			pop := heap.Pop(&hp).(pair)
			mx = max(mx, pop.value)
		}
		ans = max(ans, v+mx)
		heap.Push(&hp, pair{start: s, end: e, value: v})
	}
	return ans
}

type pair struct {
	start int
	end   int
	value int
}

type MinHeap []pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].end <= h[j].end }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
