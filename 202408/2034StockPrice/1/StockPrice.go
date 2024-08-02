package main

import (
	"container/heap"
)

func main() {

}

type pair struct {
	va int
	ti int
}

type StockPrice struct {
	Cnt  map[int]*pair
	Data []*pair
	Mx   MaxHeap
	Mi   MinHeap
}

func Constructor() StockPrice {
	return StockPrice{
		Cnt:  map[int]*pair{},
		Data: make([]*pair, 0),
		Mx:   make(MaxHeap, 0),
		Mi:   make(MinHeap, 0),
	}
}

func (this *StockPrice) Update(timestamp int, price int) {

	p := this.Cnt[timestamp]
	if p != nil {
		p.va = price
		heap.Init(&this.Mx)
		heap.Init(&this.Mi)
		return
	}

	th := &pair{ti: price, va: price}
	this.Cnt[timestamp] = th
	this.Data = append(this.Data, th)

	heap.Push(&this.Mx, th)
	heap.Push(&this.Mi, th)
}

func (this *StockPrice) Current() int {
	if len(this.Data) == 0 {
		return 0
	}
	return this.Data[len(this.Data)-1].va
}

func (this *StockPrice) Maximum() int {
	return this.Mx[0].va
}

func (this *StockPrice) Minimum() int {
	return this.Mi[0].va
}

type MaxHeap []*pair

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].va > h[j].va }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(*pair))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinHeap []*pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].va < h[j].va }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
