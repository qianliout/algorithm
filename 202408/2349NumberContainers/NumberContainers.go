package main

import (
	"container/heap"
)

func main() {

}

type NumberContainers struct {
	HP  map[int]MinHeap
	Idx map[int]int
}

func Constructor() NumberContainers {
	c := NumberContainers{
		HP:  make(map[int]MinHeap),
		Idx: make(map[int]int),
	}
	return c
}

func (this *NumberContainers) Change(index int, number int) {
	this.Idx[index] = number
	hm := this.HP[number]
	if len(hm) == 0 {
		hm = make(MinHeap, 0)
	}
	heap.Push(&hm, index)
	this.HP[number] = hm // 一定要重新赋值，不然就不能更新
}

func (this *NumberContainers) Find(number int) int {
	hm := this.HP[number]
	for hm.Len() > 0 {
		mi := (hm)[0]
		if this.Idx[mi] != number {
			heap.Pop(&hm)
			continue
		}
		this.HP[number] = hm // 一定要重新赋值，不然就不能更新
		return mi
	}
	return -1
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
