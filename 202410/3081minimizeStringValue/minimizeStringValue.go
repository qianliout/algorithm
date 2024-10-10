package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(minimizeStringValue("???"))
}

func minimizeStringValue(s string) string {
	fre := make([]Item, 26)
	for i := range fre {
		fre[i].C = byte(i + 'a')
	}

	for _, ch := range s {
		if ch == '?' {
			continue
		}
		idx := int(ch - 'a')
		fre[idx].P++
	}
	hp := make(MinHeap, 0)
	for i := range fre {
		heap.Push(&hp, fre[i])
	}
	cnt := strings.Count(s, "?")
	lost := make([]byte, 0)
	for cnt > 0 {
		pop := heap.Pop(&hp).(Item)
		lost = append(lost, pop.C)
		pop.P++
		heap.Push(&hp, pop)
		cnt--
	}
	ss := []byte(s)
	sort.Slice(lost, func(i, j int) bool { return lost[i] < lost[j] })
	j := 0
	for i, ch := range ss {
		if ch == '?' {
			ss[i] = lost[j]
			j++
		}
	}
	return string(ss)

}

type Item struct {
	C byte
	P int
}

type MinHeap []Item

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].P != h[j].P {
		return h[i].P < h[j].P
	}
	return h[i].C <= h[j].C
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
